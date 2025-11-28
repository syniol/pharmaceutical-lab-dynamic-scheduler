package lab

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Lab struct {
	Equipments []*Equipment
}

// NewLabSimulator Simulates the lab equipment and workflows
// lab is not used now, but it could be used to match IDs in workflow instead of reference
// workflows are set of tasks
func NewLabSimulator(lab *Lab, workflows ...*Workflow) error {
	var wait sync.WaitGroup
	wait.Add(len(workflows))

	for _, workflow := range workflows {
		go func() {
			done := false
			for done == false {
				doneCount := 0

				for _, task := range workflow.Tasks {
					if task.TaskStatus == TaskStatusNew {
						if task.PreviousNode != nil && task.PreviousNode.TaskStatus != TaskStatusCompleted {
							break
						}

						// todo: include the Duration check for previous and after task for expiration

						if task.equipment.status == EquipmentStatusIdle {
							task.Start()
							time.Sleep(task.Duration)
							task.Done()

							continue
						}
					}

					if task.TaskStatus == TaskStatusCompleted {
						if doneCount == len(workflow.Tasks)-1 {
							log.Println(fmt.Sprintf("Finished every task in the workflow"))

							done = true
							wait.Done()
							break
						}
						doneCount++
					}
				}
			}
		}()
	}
	wait.Wait()

	return nil
}
