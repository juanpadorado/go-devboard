// Package domain defines the TaskStatus type and its associated methods for validation and JSON unmarshalling.
package domain

import (
	"encoding/json"
	"fmt"
)

// TaskStatus represents the status of a task in the system.
type TaskStatus string

// Predefined valid task statuses.
const (
	TaskStatusTodo       TaskStatus = "TODO"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusDone       TaskStatus = "DONE"
)

// IsValid checks if the TaskStatus is one of the predefined valid statuses.
func (ts TaskStatus) IsValid() bool {
	switch ts {
	case TaskStatusTodo, TaskStatusInProgress, TaskStatusDone:
		return true
	default:
		return false
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface for TaskStatus.
func (ts *TaskStatus) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	statusTask := TaskStatus(raw)
	if !statusTask.IsValid() {
		return fmt.Errorf("invalid task status: %s", raw)
	}

	*ts = statusTask
	return nil
}
