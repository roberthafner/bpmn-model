package model

import (
	"reflect"
	"testing"
)

func TestBpmnModel_Add(t *testing.T) {
	model := BpmnModel{}

	if 0 != len(model.Processes) {
		t.Fail()
	}
	model.Add(&Process{})

	if 1 != len(model.Processes) {
		t.Fail()
	}
}

func TestBpmnModel_CurrentProcess(t *testing.T) {
	model := BpmnModel{}
	p := &Process{}
	model.Add(p)

	if 1 != len(model.Processes) {
		t.Fail()
	}

	cp := model.CurrentProcess()
	if ! reflect.DeepEqual(*p, *cp) {
		t.Fail()
	}
}

//func TestBpmnModel_SetCurrentProcess(t *testing.T) {
//	model := BpmnModel{}
//
//	if 0 != len(model.Processes) {
//		t.Fail()
//	}
//
//	p1 := &Process{}
//	model.SetCurrentProcess(p1)
//	cp := model.CurrentProcess()
//	if p1 != cp {
//		t.Fail()
//	}
//
//	p2 := &Process{}
//	model.SetCurrentProcess(p2)
//	cp = model.CurrentProcess()
//	if p2 != cp {
//		t.Fail()
//	}
//}

func TestBpmnModel_GetFlowElement(t *testing.T) {
	model := BpmnModel{}
	p := &Process{}
	model.Add(p)

	p.Add(NewUserTask("123", "task", ""))
	task := p.GetFlowElement("123")
	if "123" != task.Id() {
		t.Fail()
	}
}