package data

type TaskStatus string

const (
	Assigned   TaskStatus = "Назначена"
	InProgress TaskStatus = "В работе"
	Completed  TaskStatus = "Выполненный"
	Rejected   TaskStatus = "Отклонена"
)
