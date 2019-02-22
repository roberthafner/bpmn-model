package model

type FlowNode interface {
	FlowElement
	Asynchronous() bool
	NotExclusive() bool
	IncomingFlows() []SequenceFlow
	OutgoingFlows() []SequenceFlow

	// TODO: Methods
	// Behavior()
}

type flowNode struct {
	flowElement
	asynchronous bool
	notExclusive bool
	incoming     []SequenceFlow
	outgoing     []SequenceFlow

	// TODO: Fields?
	// behavior
}

func (fn flowNode) Asynchronous() bool {
	return fn.asynchronous
}

func (fn flowNode) NotExclusive() bool {
	return fn.notExclusive
}

func (fn flowNode) IncomingFlows() []SequenceFlow {
	return fn.incoming
}

func (fn flowNode) OutgoingFlows() []SequenceFlow {
	return fn.outgoing
}
