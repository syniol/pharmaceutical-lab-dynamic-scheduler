package lab

import "fmt"

type Workflow struct {
	ID    string
	Tasks []*Task
}

func NewWorkflow(ID string, tasks ...*Task) (*Workflow, error) {
	if len(tasks) < 2 {
		return &Workflow{
			ID:    ID,
			Tasks: tasks,
		}, nil
	}

	for i, task := range tasks {
		if i > 2 {
			task.PreviousNode = tasks[i-1]
		}
		if i < len(tasks) && i != len(tasks)-1 {
			fmt.Println("before last last", task.equipment.Name, i)
			task.NextNode = tasks[i+1]
		}

		if i == len(tasks)-1 {
		}
	}

	return &Workflow{
		ID:    ID,
		Tasks: tasks,
	}, nil
}
