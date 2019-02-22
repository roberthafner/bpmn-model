package model

type EndEvent struct {
	catchEvent
	// TODO: Methods
}

func NewEndEvent(id string, name string, documentation string) EndEvent {
	ee := EndEvent{
		catchEvent: catchEvent{
			event: event{
				flowNode: flowNode{
					flowElement: flowElement{
						baseElement: baseElement{
							id: id,
						},
						name:          "name",
						documentation: "documentation",
					},
				},
			},
		},
	}
	return ee
}
