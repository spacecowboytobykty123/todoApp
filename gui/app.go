package main

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"todoapp/backend/pkg/data"
	"todoapp/backend/pkg/jsonlog"
)

type App struct {
	logger *jsonlog.Logger
	models data.Models
}

func (a *App) startup(ctx context.Context) {
	a.logger.PrintInfo("GUI started", nil)
}

// NewApp инициализация приложения с подключением к БД
func NewApp() *App {
	logger := jsonlog.New(os.Stdout, jsonlog.LeverInfo)

	dsn := "postgres://crm:pass@localhost/tasks?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(15 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}

	return &App{
		logger: logger,
		models: data.NewModels(db),
	}
}

//
// ========== Методы, доступные во фронтенде (Wails Bind) ==========
//

// GetTasks — вернуть список задач
func (a *App) GetTasks(status string, deadline string) ([]*data.Task, error) {
	var taskStatus *data.TaskStatus
	if status != "" {
		ts := data.TaskStatus(status)
		taskStatus = &ts
	}

	var endDate *time.Time
	if deadline != "" {
		t, err := time.Parse("2006-01-02", deadline) // 👈 формат из <input type="date">
		if err == nil {
			endDate = &t
		}
	}

	// тут можно прокинуть фильтры (например, страница = 1, размер = 50)
	filters := data.Filters{Page: 1, PageSize: 50}

	tasks, _, err := a.models.Tasks.GetAllTasks(nil, endDate, taskStatus, filters)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTask — получить задачу по ID
func (a *App) GetTask(id string) (*data.Task, error) {
	taskID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return a.models.Tasks.GetTask(taskID)
}

// CreateTask — создать новую задачу
func (a *App) CreateTask(name, desc string, status data.TaskStatus, deadline *time.Time) (*data.Task, error) {
	task := &data.Task{
		Name:     name,
		Desc:     desc,
		Status:   status,
		Deadline: deadline,
	}
	err := a.models.Tasks.InsertTask(task)
	return task, err
}

// UpdateTask — обновить задачу
func (a *App) UpdateTask(id string, name, desc *string, status *data.TaskStatus, deadline *time.Time) (*data.Task, error) {
	taskID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	task, err := a.models.Tasks.GetTask(taskID)
	if err != nil {
		return nil, err
	}

	if name != nil {
		task.Name = *name
	}
	if desc != nil {
		task.Desc = *desc
	}
	if status != nil {
		task.Status = *status
	}
	if deadline != nil {
		task.Deadline = deadline
	}

	if err := a.models.Tasks.UpdateSubscription(task); err != nil {
		return nil, err
	}

	return task, nil
}

// DeleteTask — удалить задачу
func (a *App) DeleteTask(id string) error {
	taskID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return a.models.Tasks.DeleteTask(taskID)
}
