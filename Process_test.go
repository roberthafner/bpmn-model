package model

import "testing"

func TestProcess(t *testing.T) {
	p := &Process {
		baseElement: baseElement{
			id:                  "123",
			xmlRowNumber:        30,
			xmlColumnNumber:     20,
			extensionElements:   map[string][]ExtensionElement{"ee": []ExtensionElement{ExtensionElement{}}},
			extensionAttributes: map[string][]ExtensionAttribute{"ea": []ExtensionAttribute{ExtensionAttribute{}}},
		},
		Name: "name",
		Executable: true,
		Documentation: "documentation",
	}

	if "123" != p.Id() {
		t.Fail()
	}

	if 30 != p.xmlRowNumber {
		t.Fail()
	}

	if 20 != p.xmlColumnNumber {
		t.Fail()
	}

	if 1 != len(p.extensionElements) {
		t.Fail()
	}

	if 1 != len(p.extensionAttributes) {
		t.Fail()
	}

	if "name" != p.Name {
		t.Fail()
	}

	if "documentation" != p.Documentation {
		t.Fail()
	}
}

func TestProcess_Add(t *testing.T) {
	p := &Process{
		baseElement: baseElement{
			id:                  "123",
			xmlRowNumber:        30,
			xmlColumnNumber:     20,
			extensionElements:   map[string][]ExtensionElement{"ee": []ExtensionElement{ExtensionElement{}}},
			extensionAttributes: map[string][]ExtensionAttribute{"ea": []ExtensionAttribute{ExtensionAttribute{}}},
		},
		Name:          "name",
		Executable:    true,
		Documentation: "documentation",
	}

	if 0 != len(p.FlowElements) {
		t.Fail()
	}

	fe := flowElement{}
	p.Add(fe)


	if 1 != len(p.FlowElements) {
		t.Fail()
	}
}

func TestProcess_GetFlowElement(t *testing.T) {
	p := &Process{
		baseElement: baseElement{
			id:                  "123",
			xmlRowNumber:        30,
			xmlColumnNumber:     20,
			extensionElements:   map[string][]ExtensionElement{"ee": []ExtensionElement{ExtensionElement{}}},
			extensionAttributes: map[string][]ExtensionAttribute{"ea": []ExtensionAttribute{ExtensionAttribute{}}},
		},
		Name:          "name",
		Executable:    true,
		Documentation: "documentation",
	}

	fe := flowElement{
		baseElement: baseElement {id: "123",},
	}
	p.Add(fe)


	fe2 := p.GetFlowElement("123")
	if fe2.Id() != fe.Id() {
		t.Fail()
	}
}

func TestProcess_NewProcess(t *testing.T) {
	p := NewProcess("123", "name", "doc")

	if "123" != p.Id() {
		t.Fail()
	}

	if "name" != p.Name {
		t.Fail()
	}

	if "doc" != p.Documentation {
		t.Fail()
	}
}

func TestProcess_String(t *testing.T) {
	p := &Process{
		baseElement: baseElement{
			id:                  "123",
			xmlRowNumber:        30,
			xmlColumnNumber:     20,
			extensionElements:   map[string][]ExtensionElement{"ee": []ExtensionElement{ExtensionElement{}}},
			extensionAttributes: map[string][]ExtensionAttribute{"ea": []ExtensionAttribute{ExtensionAttribute{}}},
		},
		Name:          "name",
		Executable:    true,
		Documentation: "documentation",
	}

	s := p.String()
	if "Process: (id=123, name=name)\n" != s {
		t.Fail()
	}
}

// Verify flowElement implements BaseElement
var _ BaseElement = (*flowElement)(nil)
