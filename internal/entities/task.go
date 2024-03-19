package entities

type Task struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	Creator     int64  `json:"creator_id"`
}

func (t *Task) SetPriorityLow() {
	t.Priority = "low"
}

func (t *Task) SetPriorityMedium() {
	t.Priority = "medium"
}

func (t *Task) SetPriorityHigh() {
	t.Priority = "high"
}

func (t *Task) SetStatusInProgress() {
	t.Status = "in progress"
}

func (t *Task) SetStatusPending() {
	t.Priority = "pending"
}

func (t *Task) SetStatusCompleted() {
	t.Priority = "completed"
}
