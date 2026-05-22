package types

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"strings"
)

type Task struct {
	ID            string       `json:"id" db:"id"`
	ProjectID     string       `json:"projectID" db:"project_id"`
	Title         string       `json:"title" db:"title"`
	IsAllDay      bool         `json:"isAllDay" db:"is_all_day"`
	DueDate       TuidoTime    `json:"dueDate,omitzero" db:"due_date,omitzero"`
	CompletedTime TuidoTime    `json:"completedTime,omitzero" db:"completed_time,omitzero"`
	Description   string       `json:"description" db:"description"`
	Status        Status       `json:"status" db:"status"`
	Priority      TaskPriority `json:"priority" db:"priority"`
	SortOrder     int64        `json:"sortOrder" db:"sort_order"`
	Created_at    TuidoTime    `json:"createdAt" db:"created_at"`
	Deleted_at    TuidoTime    `json:"deletedAt" db:"deleted_at"`
	Updated_at    TuidoTime    `json:"updatedAt" db:"updated_at"`
}

type Status int

var (
	StatusNormal   Status = 0
	StatusComplete Status = 2
)

func (s *Status) UnmarshalJSON(data []byte) error {
	var status int
	if err := json.Unmarshal(data, &status); err != nil {
		return err
	}
	switch status {
	case int(StatusNormal), int(StatusComplete):
		*s = Status(status)
	default:
		*s = StatusNormal
	}
	return nil
}

func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(s))
}

func (s Status) String() string {
	switch s {
	case StatusComplete:
		return color.Green.Sprint("☑")
	case StatusNormal:
		return color.White.Sprint("☐")
	default:
		return color.Red.Sprint("☒")
	}
}

type TaskPriority int

const (
	TaskPriorityNone   TaskPriority = 0
	TaskPriorityLow    TaskPriority = 1
	TaskPriorityMedium TaskPriority = 3
	TaskPriorityHigh   TaskPriority = 5
)

var (
	NonePriorityColor   = color.HEX("#C6C6C6").C256()
	LowPriorityColor    = color.HEX("#4772F9").C256()
	MediumPriorityColor = color.HEX("#FAA80B").C256()
	HighPriorityColor   = color.HEX("#D52B24").C256()
)

var priorityMap = map[string]TaskPriority{
	"none":   TaskPriorityNone,
	"low":    TaskPriorityLow,
	"medium": TaskPriorityMedium,
	"high":   TaskPriorityHigh,
}

func (p *TaskPriority) UnmarshalJSON(data []byte) error {
	var priority int
	if err := json.Unmarshal(data, &priority); err != nil {
		return err
	}
	switch priority {
	case int(TaskPriorityNone), int(TaskPriorityLow), int(TaskPriorityMedium), int(TaskPriorityHigh):
		*p = TaskPriority(priority)
	default:
		*p = TaskPriorityNone
	}
	return nil
}

func (p TaskPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(p))
}

func (p TaskPriority) String() string {
	flag := "⚑"
	switch p {
	case TaskPriorityNone:
		flag = NonePriorityColor.Sprint(flag)
	case TaskPriorityLow:
		flag = LowPriorityColor.Sprint(flag)
	case TaskPriorityMedium:
		flag = MediumPriorityColor.Sprint(flag)
	case TaskPriorityHigh:
		flag = HighPriorityColor.Sprint(flag)
	}

	return flag
}

func (p *TaskPriority) Set(value string) error {
	priority, ok := priorityMap[strings.ToLower(value)]
	if !ok {
		return fmt.Errorf("invalid task: %s", value)
	}

	*p = priority
	return nil
}
