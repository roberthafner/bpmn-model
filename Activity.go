package model

type Activity interface {
	FlowNode

	// TODO: Methods
	// DefaultFlow() string
	// ForCompensation() bool
	// LoopCharacteristics() MultiInstanceLoopCharacteristics
	// IoSpecification() IOSpecification
	// DataInputAssociations() []DataAssociation
	// DataOutputAssociations() []DataAssociation
	// FailedJobRetryTimeCycleValue() string
	// MapExceptions() []MaxExceptionEntry
}

type activity struct {
	flowNode

	// TODO: Fields
	// defaultFlow string
	// forCompensation bool
	// loopCharacteristics MultiInstanceLoopCharacteristics
	// ioSpecification IOSpecification
	// dataInputAssociations []DataAssociation
	// dataOutputAssociations []DataAssociation
	// failedJobRetryTimeCycleValue string
	// mapExceptions []MaxExceptionEntry

	// TODO: Methods (for fields)
}
