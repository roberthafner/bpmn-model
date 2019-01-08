package model

import "testing"

func TestDocumentation_Fields(t *testing.T) {
	d := Documentation {
		Text : "doc",
		TextFormat : "text/plain",
	}

	if "doc" != d.Text {
		t.Fail()
	}

	if "text/plain" != d.TextFormat {
		t.Fail()
	}
}

// Verify baseElement implements BaseElement
var _ BaseElement = (*baseElement)(nil)

// Verify flowElement implements FlowElement
var _ FlowElement = (*flowElement)(nil)

// Verify flowNode implements FlowNode
var _ FlowNode = (*flowNode)(nil)
