package model

import "strings"

type BpmnModel struct {
	//DefinitionAttributes map[string][]ExtensionAttribute
	Processes []Process
	// LocationMap map[string]GraphicInfo
	// LabelLocationMap map[string]GraphicInfo
	// FlowLocationMap map[string]GraphicInfo
	// Signals []Signal
	// MessageFlowMap map[string]MessageFlow
	// MessageMap map[string]Message
	// ErrorMap map[string]string
	// ItemDefinitionMap map[string]ItemDefinition
	// DataStoreMap map[string]DataStore
	// Pools []Pool
	// Imports []Import
	// Interfaces []Interface
	// Artifacts []Artifact
	// Resource []Resources
	// NamespaceMap map[string]string
	// TargetNamespace string
	// SourceSystemId string
	// UserTaskFormTypes []string
	// StartEventFormTypes []string
	// NextFlowIdCounter int
	// EventSupport

	// TODO: Fields
	// TODO: Methods
}

func (model *BpmnModel) Add(process *Process) {
	if nil == model.Processes {
		model.Processes = make([]Process, 0, 10)
	}
	model.Processes = append(model.Processes, *process)
}

func (model *BpmnModel) CurrentProcess() *Process {
	return &model.Processes[len(model.Processes)-1]
}

//func (model *BpmnModel) SetCurrentProcess(p *Process) {
//	len := len(model.Processes)
//	if len != 0 {
//		len = len -1
//	}
//	model.Processes[len] = *p
//}

func (model *BpmnModel) GetFlowElement(flowElementId string) FlowElement {
	for i := 0; i < len(model.Processes); i++ {
		var flowElement = model.Processes[i].GetFlowElement(flowElementId)
		if nil != flowElement {
			return flowElement
		}
	}
	return nil
}

func (model *BpmnModel) String() string {
	var s strings.Builder
	for i := 0; i < len(model.Processes); i++ {
		s.WriteString(model.Processes[i].String())
		s.WriteString("\n")
	}
	return s.String()
}
