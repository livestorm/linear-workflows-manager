package linear

type BoardState struct {
	Id   string
	Name string
}

type ColumnMapping struct {
	ChildColumn  string
	ParentColumn string
	Comment      string
}

type BoardConfig struct {
	BoardStates     []BoardState
	SquadTribes     map[string]string
	TribeChannels   map[string]string
	MandatoryLabels map[string]map[string][]string
	ColumnMappings  map[string][]ColumnMapping
	DefaultStateId  string
}

type TicketConfig struct {
	Headers []string
}
