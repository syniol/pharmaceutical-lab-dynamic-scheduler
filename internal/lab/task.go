package lab

import (
	"time"

	"automata/pkg"
)

type TaskStatus string

const (
	TaskStatusNew       TaskStatus = "NOT_STARTED"
	TaskStatusInProcess TaskStatus = "IN_PROCESS"
	TaskStatusCompleted TaskStatus = "COMPLETED"
	TaskStatusFailed    TaskStatus = "FAILED"
)

type Task struct {
	*pkg.Node[Task]
	equipment *Equipment

	Command            string
	Duration           time.Duration
	ExpirationDuration time.Duration
	TaskStatus         TaskStatus
	startDateTime      time.Time
}

func NewTask(equipment *Equipment) *Task {
	return &Task{
		TaskStatus: TaskStatusNew,
		equipment:  equipment,
		Node:       &pkg.Node[Task]{},
	}
}

func (task *Task) Start() {
	task.equipment.status = EquipmentStatusBusy
	task.startDateTime = time.Now()
	task.TaskStatus = TaskStatusInProcess
}

func (task *Task) Done() {
	task.equipment.status = EquipmentStatusIdle
	task.TaskStatus = TaskStatusCompleted
}

func (task *Task) FinishDateTime() time.Time {
	startDateTime := task.startDateTime
	return startDateTime.Add(task.Duration)
}
