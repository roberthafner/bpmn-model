package model

type ExtensionElement struct {
	baseElement
	Name            string
	NamespacePrefix string
	Namespace       string
	ElementText     string
	ChildElements   []ExtensionElement
	// TODO: Methods?
}
