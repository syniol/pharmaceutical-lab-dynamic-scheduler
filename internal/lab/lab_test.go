package lab

import "testing"

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

	workflowTwo, _ := NewWorkflow(
		"B",
		NewTask(incubator),
		NewTask(extractor),
		NewTask(analyzer),
	)

	actual := NewLabSimulator(lab, workflow, workflowTwo)
	if actual != nil {
		t.Error("expected no error but got, ", actual)
	}
}
