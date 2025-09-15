package data

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"time"
	"todoapp/backend/pkg/validator"
)

type Task struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Status    TaskStatus `json:"status"`
	Version   int        `json:"version"`
	CreatedAt time.Time  `json:"created_at"`
	Deadline  *time.Time `json:"deadline"`
}

type TaskModel struct {
	DB *sql.DB
}

func ValidateTask(v *validator.Validator, task *Task) {
	v.Check(task.Name != "", "name", "нужно ввести имя задачи!")
}

func (t TaskModel) InsertTask(task *Task) error {
	query := `INSERT INTO tasks (name, description, status, deadline)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, version`

	args := []any{task.Name, task.Desc, task.Status, task.Deadline}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return t.DB.QueryRowContext(ctx, query, args...).Scan(&task.ID, &task.CreatedAt, &task.Version)
}

func (t TaskModel) GetTask(id uuid.UUID) (*Task, error) {
	query := `SELECT id, name, description, status, created_at, version, deadline
FROM tasks
WHERE id = $1`

	var task Task

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := t.DB.QueryRowContext(ctx, query, id).Scan(
		&task.ID,
		&task.Name,
		&task.Desc,
		&task.Status,
		&task.CreatedAt,
		&task.Version,
		&task.Deadline,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &task, err
}

func (t TaskModel) UpdateSubscription(task *Task) error {
	query := `UPDATE tasks
	SET name = $1, description = $2, status = $3, deadline = $4, version = version + 1
	WHERE id = $5 and version = $6
	RETURNING version
`

	args := []any{task.Name, task.Desc, task.Status, task.Deadline, task.ID, task.Version}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := t.DB.QueryRowContext(ctx, query, args...).Scan(&task.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err

		}
	}
	return nil
}

func (t TaskModel) DeleteTask(id uuid.UUID) error {
	query := `DELETE FROM tasks WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := t.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (t TaskModel) GetAllTasks(startDate, endDate *time.Time, taskStatus *TaskStatus, filters Filters) ([]*Task, Metadata, error) {
	query := `SELECT COUNT(*) OVER() AS total_count, id, name, description, status, created_at, version, deadline 
FROM tasks
WHERE ($1::task_status IS NULL OR status = $1::task_status)
AND ($2::date IS NULL OR created_at >= $2::date)
AND ($3::date IS NULL OR deadline <= $3::date)
ORDER BY created_at DESC
LIMIT $4 OFFSET $5;
`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := t.DB.QueryContext(ctx, query, taskStatus, startDate, endDate, filters.limit(), filters.offset())
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	tasks := []*Task{}

	for rows.Next() {
		var task Task

		err := rows.Scan(
			&totalRecords,
			&task.ID,
			&task.Name,
			&task.Desc,
			&task.Status,
			&task.CreatedAt,
			&task.Version,
			&task.Deadline,
		)

		if err != nil {
			return nil, Metadata{}, err
		}

		tasks = append(tasks, &task)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return tasks, metadata, err
}
