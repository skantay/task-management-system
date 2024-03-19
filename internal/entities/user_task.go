package entities

type UserTask struct {
	UserID int64 `json:"user_id" db:"user_id"`
	TaskID int64 `json:"task_id" db:"task_id"`
}
