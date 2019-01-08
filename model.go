package model

import "strings"

const (
	AttributeId string = "id"
	AttributeName string = "name"

	AttributeSourceRef= "sourceRef"
	AttributeTargetRef= "targetRef"

	ElementProcess string = "process"
	ElementDocumentation string = "documentation"
	ElementStartEvent string = "startEvent"
	ElementEndEvent string = "endEvent"
	ElementUserTask string = "userTask"
	ElementSequenceFlow string = "sequenceFlow"
)

type Documentation struct {
	Text string
	TextFormat string
}

type ExtensionElement struct {
	baseElement
	Name string
	NamespacePrefix string
	Namespace string
	ElementText string
	ChildElements []ExtensionElement
	// TODO: Methods?
}

type ExtensionAttribute struct {
	NamespacePrefix string
	Namespace string
	Name string
	Value string
	// TODO: Methods?
}

func (ea ExtensionAttribute) String() string {
	var s strings.Builder
	if len(ea.NamespacePrefix) == 0 {
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

type BaseElement interface {
	Id() string
	ExtensionElements() map[string][]ExtensionElement
	ExtensionAttributes() map[string][]ExtensionAttribute
	// TODO:
	//XmlRowNumber() uint
	//XmlColumnNumber() uint
}

type FlowElement interface {
	BaseElement
	Name() string

	// TODO:
	/*
	protected List<FlowableListener> executionListeners = new ArrayList<>();
	protected FlowElementsContainer parentContainer;
	*/

}

type SequenceFlow struct {
	flowElement
	SourceRef string
	TargetRef string
	ConditionExpression string
	SourceFlowElement FlowElement
	TargetFlowElement FlowElement
}

type FlowNode interface {
	FlowElement
	IncomingFlows() []SequenceFlow
	OutgoingFlows() []SequenceFlow

	// TODO:
	/*
	protected boolean asynchronous;
	protected boolean notExclusive;
	*/

}

type FlowElementsContainer interface {
	FlowElement() FlowElement
	FlowElements() []FlowElement
	FlowElementMap() map[string]FlowElement

	// TODO:
	/*
	void addFlowElement(FlowElement element);
	void addFlowElementToMap(FlowElement element);
	void removeFlowElement(String elementId);
	void removeFlowElementFromMap(String elementId);
	Artifact getArtifact(String id);
	Collection<Artifact> getArtifacts();
	void addArtifact(Artifact artifact);
	void removeArtifact(String artifactId);
	*/
}

type Activity interface {
	FlowNode
	// TODO:
}

type Task interface {
	Activity
	// TODO:
}

type Event interface {
	FlowNode
	// TODO:
}

type ThrowEvent interface {
	Event
	// TODO:
}

type CatchEvent interface {
	Event
	// TODO: other fields
}

type StartEvent struct {
	event
	// TODO:
}

type EndEvent struct {
	event
	// TODO:
}

type UserTask struct {
	task
	assignee string
	candidateUsers []string
	candidateGroups []string
	actualOwner string
	// TODO:
}

type Process struct {
	baseElement
	Name string
	FlowElements []FlowElement
	FlowElementMap map[string]FlowElement
	// TODO:
}

func (p Process) Add(element FlowElement) Process {
	if p.FlowElements == nil {
		p.FlowElements = make([]FlowElement, 0, 10)
		p.FlowElementMap = make(map[string]FlowElement, 10)
	}
	p.FlowElements = append(p.FlowElements, element)
	p.FlowElementMap[element.Id()] = element
	return p
}

func (p Process) GetFlowElement(flowElementId string) FlowElement {
	return p.FlowElementMap[flowElementId]
}

func (p Process) String() string {
	var s strings.Builder
	s.WriteString("Process: (id=")
	s.WriteString(p.Id())
	s.WriteString(", name=")
	s.WriteString(p.Name)
	s.WriteString(")\n")

	for _, fe := range p.FlowElements {
		s.WriteString(String(fe))
	}

	return s.String()
}

func String(fe FlowElement) string {
	var s strings.Builder
	s.WriteString("FlowElement: (id=")
	s.WriteString(fe.Id())
	s.WriteString(", name=")
	s.WriteString(fe.Name())
	s.WriteString(")\n")

	return s.String()
}

type BpmnModel struct {
	processes []Process
}

func (model BpmnModel) Add(process Process) BpmnModel{
	if nil == model.processes {
		model.processes = make([]Process,0,10)
	}
	model.processes = append(model.processes, process)
	return model
}

func (model BpmnModel) Processes() []Process {
	return model.processes
}

func (model BpmnModel) CurrentProcess() Process {
	return model.processes[len(model.Processes())-1]
}

func (model BpmnModel) SetCurrentProcess(p Process) {
	model.processes[len(model.Processes())-1] = p
}

func (model BpmnModel) GetFlowElement(flowElementId string) FlowElement {
	for i := 0; i < len(model.processes); i++ {
		var flowElement = model.processes[i].GetFlowElement(flowElementId)
		if nil != flowElement {
			return flowElement
		}
	}
	return nil
}

func (model BpmnModel) String() string {
	var s strings.Builder
	for i := 0; i < len(model.processes); i++ {
		s.WriteString(model.processes[i].String())
		s.WriteString("\n")
	}
	return s.String()
}

type baseElement struct {
	id string
	documentation []Documentation
	extensionElements map[string][]ExtensionElement
	extensionAttributes map[string][]ExtensionAttribute
}

func (be baseElement) Id() string {
	return be.id
}

func (be baseElement) Documentation() []Documentation {
	return be.documentation
}

func (be baseElement) ExtensionElements() map[string][]ExtensionElement {
	return be.extensionElements
}

func (be baseElement) ExtensionAttributes() map[string][]ExtensionAttribute {
	return be.extensionAttributes
}

type flowElement struct {
	baseElement
	name string
}

func (fe flowElement) Name() string {
	return fe.name
}

type flowNode struct {
	flowElement
	incoming []SequenceFlow
	outgoing []SequenceFlow
}

func (fn flowNode) IncomingFlows() []SequenceFlow {
	return fn.incoming
}

func (fn flowNode) OutgoingFlows() []SequenceFlow {
	return fn.outgoing
}

type event struct {
	flowNode
}

type activity struct {
	flowNode
}

type task struct {
	activity
}

func NewEndEvent(id string, name string, documentation []Documentation) EndEvent {
	ee := EndEvent{}
	ee.id = id
	ee.name = name
	ee.documentation = documentation
	return ee
}

func NewProcess(id string, name string, documentation []Documentation) Process {
	p := Process {}
	p.id = id
	p.documentation = documentation
	p.Name = name
	return p
}

func NewSequenceFlow(id string, name string, documentation []Documentation, sourceRef string, targetRef string) SequenceFlow {
	sf := SequenceFlow{}
	sf.id = id
	sf.documentation = documentation
	sf.name = name
	sf.SourceRef = sourceRef
	sf.TargetRef = targetRef
	return sf
}

func NewStartEvent(id string, name string, documentation []Documentation) StartEvent {
	se := StartEvent{}
	se.id = id
	se.name = name
	se.documentation = documentation
	return se
}

func NewUserTask(id string, name string, documentation []Documentation) UserTask {
	ut := UserTask{}
	ut.id = id
	ut.name = name
	ut.documentation = documentation
	return ut
}
