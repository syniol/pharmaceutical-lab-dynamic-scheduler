package lab

import (
	"automata/pkg"
)

type EquipmentStatus string

const (
	EquipmentStatusIdle EquipmentStatus = "IDLE"
	EquipmentStatusBusy EquipmentStatus = "BUSY"
)

type Equipment struct {
	*pkg.Node[Task]

	Name   string
	status EquipmentStatus
}

func NewEquipment(name string) *Equipment {
	return &Equipment{
		Name:   name,
		status: EquipmentStatusIdle,
	}
}

func (e *Equipment) SetStatus(status EquipmentStatus) {
	e.status = status
}
