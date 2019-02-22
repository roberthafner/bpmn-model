package model

type BaseElement interface {
	Id() string
	XmlRowNumber() uint
	XmlColumnNumber() uint
	ExtensionElements() map[string][]ExtensionElement
	ExtensionAttributes() map[string][]ExtensionAttribute
	// TODO: Methods?

}

type baseElement struct {
	id                  string
	xmlRowNumber        uint
	xmlColumnNumber     uint
	extensionElements   map[string][]ExtensionElement
	extensionAttributes map[string][]ExtensionAttribute
}

func (be baseElement) Id() string {
	return be.id
}

func (be baseElement) XmlRowNumber() uint {
	return be.xmlRowNumber
}

func (be baseElement) XmlColumnNumber() uint {
	return be.xmlColumnNumber
}

func (be baseElement) ExtensionElements() map[string][]ExtensionElement {
	return be.extensionElements
}

func (be baseElement) ExtensionAttributes() map[string][]ExtensionAttribute {
	return be.extensionAttributes
}
