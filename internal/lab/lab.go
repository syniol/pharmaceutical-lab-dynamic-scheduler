package lab

import (
	"fmt"
	"sync"
	"time"
)

type Lab struct {
	Equipments []*Equipment
}

func NewLabSimulator(lab *Lab, workflows ...*Workflow) {
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

						if task.equipment.status == EquipmentStatusIdle {
							task.Start()
							time.Sleep(task.Duration)
							task.Done()

							continue
						}
					}

					if task.TaskStatus == TaskStatusCompleted {
						if doneCount == len(workflow.Tasks)-1 {
							fmt.Println(fmt.Sprintf("Finished all Tasks"))

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
}
