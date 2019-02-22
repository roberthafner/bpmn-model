package model


import "testing"

func TestExtensionAttribute(t *testing.T) {

	ea := ExtensionAttribute{
		Name: "Name",
		Value : "Value",
		NamespacePrefix: "NamespacePrefix",
		Namespace: "Namespace",
	}

	s := ea.String()
	if ("NamespacePrefix:Name=Value" != s) {
		t.Fail()
	}
}