package model


import "testing"

func TestEventDefinition(t *testing.T) {

	ed := eventDefinition{
		baseElement: baseElement{
			id:                  "123",
			xmlRowNumber:        30,
			xmlColumnNumber:     20,
			extensionElements:   map[string][]ExtensionElement{"ee": []ExtensionElement{ExtensionElement{}}},
			extensionAttributes: map[string][]ExtensionAttribute{"ea": []ExtensionAttribute{ExtensionAttribute{}}},
		},
	}

	if "123" != ed.Id() {
		t.Fail()
	}

	if 30 != ed.xmlRowNumber {
		t.Fail()
	}

	if 20 != ed.xmlColumnNumber {
		t.Fail()
	}

	if 1 != len(ed.extensionElements) {
		t.Fail()
	}

	if 1 != len(ed.extensionAttributes) {
		t.Fail()
	}
}