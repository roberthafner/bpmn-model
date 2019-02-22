package model

import "testing"

func TestTask(t *testing.T) {

	task := task{
		activity : activity{
			flowNode: flowNode{
				flowElement: flowElement{
					baseElement: baseElement{
						id:                  "123",
						xmlRowNumber:        30,
						xmlColumnNumber:     20,
						extensionElements:   map[string][]ExtensionElement{"ee": []ExtensionElement{ExtensionElement{}}},
						extensionAttributes: map[string][]ExtensionAttribute{"ea": []ExtensionAttribute{ExtensionAttribute{}}},
					},
					name:          "name",
					documentation: "documentation",
				},
				asynchronous: true,
				notExclusive: true,
				incoming:     []SequenceFlow{SequenceFlow{}},
				outgoing:     []SequenceFlow{SequenceFlow{}},
			},
		},
	}

	if "123" != task.Id() {
		t.Fail()
	}

	if 30 != task.xmlRowNumber {
		t.Fail()
	}

	if 20 != task.xmlColumnNumber {
		t.Fail()
	}

	if 1 != len(task.extensionElements) {
		t.Fail()
	}

	if 1 != len(task.extensionAttributes) {
		t.Fail()
	}

	if "name" != task.Name() {
		t.Fail()
	}

	if "documentation" != task.Documentation() {
		t.Fail()
	}

	if true != task.Asynchronous() {
		t.Fail()
	}

	if true != task.NotExclusive() {
		t.Fail()
	}

	if 1 != len(task.IncomingFlows()) {
		t.Fail()
	}

	if 1 != len(task.OutgoingFlows()) {
		t.Fail()
	}
}

// Verify task implements Task
var _ Task = (*task)(nil)

// Verify task implements Activity
var _ Activity = (*task)(nil)

// Verify task implements FlowNode
var _ FlowNode = (*task)(nil)

// Verify task implements FlowElement
var _ FlowElement = (*task)(nil)

// Verify StartEvent implements BaseElement
var _ BaseElement = (*task)(nil)
