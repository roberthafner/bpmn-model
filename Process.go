package model

import "strings"

type Process struct {
	baseElement
	Name          string
	Executable    bool
	Documentation string
	// IoSpecification IOSpecification
	// ExecutionListeners []
	// Lanes[] []Lane
	FlowElements []FlowElement
	// DataObjects []
	// Artifacts []
	// CandidateStarterUsers []string
	// CandidateStarterGroups []string
	// EventListeners []EventListener
	FlowElementMap map[string]FlowElement

	//InitialFlowElement FlowElement
	//EnableEagerExecutionTreeFetching bool

	// TODO: Fields
	// TODO: Methods
}

func (p *Process) Add(element FlowElement) {
	if p.FlowElements == nil {
		p.FlowElements = make([]FlowElement, 0, 10)
		p.FlowElementMap = make(map[string]FlowElement, 10)
	}
	p.FlowElements = append(p.FlowElements, element)
	p.FlowElementMap[element.Id()] = element
}

func (p *Process) GetFlowElement(flowElementId string) FlowElement {
	return p.FlowElementMap[flowElementId]
}

func (p *Process) String() string {
	var s strings.Builder
	s.WriteString("Process: (id=")
	s.WriteString(p.Id())
	s.WriteString(", name=")
	s.WriteString(p.Name)
	s.WriteString(")\n")

	for _, fe := range p.FlowElements {
		s.WriteString(fe.String())
	}

	return s.String()
}

func NewProcess(id string, name string, documentation string) *Process {
	p := &Process{}
	p.id = id
	p.Documentation = documentation
	p.Name = name
	return p
}
