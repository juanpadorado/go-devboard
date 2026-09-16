package domain

import (
	"encoding/json"
	"testing"
)

type CreateTaskRequest struct {
	Title  string     `json:"title"`
	Status TaskStatus `json:"status"`
}

func TestCreateTaskRequestWithValidStatus(t *testing.T) {
	body := []byte(`{"title": "Test Task", "status": "TODO"}`)

	var req CreateTaskRequest

	err := json.Unmarshal(body, &req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if req.Status != TaskStatusTodo {
		t.Fatalf("Expected status to be %s, got %s", TaskStatusTodo, req.Status)
	}
}

func TestCreateTaskRequestWithInvalidStatus(t *testing.T) {
	body := []byte(`{"title": "Test Task", "status": "banana"}`)

	var req CreateTaskRequest

	err := json.Unmarshal(body, &req)
	if err == nil {
		t.Fatalf("esperaba error, pero recibi nil")
	}

	expected := `invalid task status: banana`

	if err.Error() != expected {
		t.Fatalf("esperaba error %q, recibi %q", expected, err.Error())
	}
}
