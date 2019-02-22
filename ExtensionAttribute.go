package model

import (
	"strings"
)

type ExtensionAttribute struct {
	Name            string
	Value           string
	NamespacePrefix string
	Namespace       string
	// TODO: Methods?
}

func (ea ExtensionAttribute) String() string {
	var s strings.Builder
	if len(ea.NamespacePrefix) != 0 {
		s.WriteString(ea.NamespacePrefix)
		s.WriteString(":")
	}
	s.WriteString(ea.Name)
	if len(ea.Value) != 0 {
		s.WriteString("=")
		s.WriteString(ea.Value)
	}
	return s.String()
}
