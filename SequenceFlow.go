package model

type SequenceFlow struct {
	flowElement
	ConditionExpression string
	SourceRef           string
	TargetRef           string

	SourceFlowElement FlowElement
	TargetFlowElement FlowElement

	// TODO: Fields
	//WayPoints []uint
}

func NewSequenceFlow(id string, name string, documentation string, sourceRef string, targetRef string) SequenceFlow {
	sf := SequenceFlow{}
	sf.id = id
	sf.documentation = documentation
	sf.name = name
	sf.SourceRef = sourceRef
	sf.TargetRef = targetRef
	return sf
}
