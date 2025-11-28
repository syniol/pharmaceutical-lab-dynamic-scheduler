package lab

import (
	"testing"
	"time"
)

func TestNewLabSimulator(t *testing.T) {
	incubator, extractor, dispenser, analyzer :=
		NewEquipment("Incubator"),
		NewEquipment("Extractor"),
		NewEquipment("Dispenser"),
		NewEquipment("Analyzer")

	lab := &Lab{
		Equipments: []*Equipment{
			incubator,
			extractor,
			dispenser,
			analyzer,
		},
	}

	workflow, _ := NewWorkflow(
		"A",
		NewTask(incubator),
		NewTask(extractor),
	)

	taskOne := NewTask(incubator)
	taskTwo := NewTask(extractor)
	taskTwo.Duration = time.Second * 4
	taskThree := NewTask(analyzer)

	workflowTwo, _ := NewWorkflow(
		"B",
		taskOne,
		taskTwo,
		taskThree,
	)

	actual := NewLabSimulator(lab, workflow, workflowTwo)
	if actual != nil {
		t.Error("expected no error but got, ", actual)
	}
}
