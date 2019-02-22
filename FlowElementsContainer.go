package model

type FlowElementsContainer interface {
	FlowElement() FlowElement
	FlowElements() []FlowElement
	FlowElementMap() map[string]FlowElement

	// TODO: Methods?
	// AddFlowElement(fe FlowElement)
	// AddFlowElementToMap(fe FlowElement)
	// RemoveFlowElement(elementId string)
	// RemoveFlowElementFromMap(elementId string)
	// Artifact(id String) Artifact
	// Artifacts() []Artifact
	// AddArtifact(a Artifact)
	// RemoveArtifact(artifactId string)
}
