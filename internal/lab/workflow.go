package lab

type Workflow struct {
	ID    string
	Tasks []*Task
}

func NewWorkflow(ID string, tasks ...*Task) (*Workflow, error) {
	for i, task := range tasks {
		if i > 1 {
			task.PreviousNode = tasks[i-1]
		}
		if i != len(tasks)-1 {
			task.NextNode = tasks[i+1]
		}
	}

	return &Workflow{
		ID:    ID,
		Tasks: tasks,
	}, nil
}
