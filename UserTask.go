package model

type UserTask struct {
	task
	Assignee string
	Owner    string
	Priority string
	// FormKey string
	DueDate              string
	BusinessCalendarName string
	Category             string
	ExtensionId          string

	CandidateUsers  []string
	CandidateGroups []string
	//FormProperties []FormProperty
	//TaskListeners []

	// TODO: Fields
	// CustomUserIdentityLinks map[string][]string
	// CustomGroupIdentityLinks map[string][]string
	// CustomProperties []CustomProperty

	// TODO: Methods
}

func NewUserTask(id string, name string, documentation string) UserTask {
	ut := UserTask{}
	ut.id = id
	ut.name = name
	ut.documentation = documentation
	return ut
}
