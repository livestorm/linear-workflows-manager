package linear

type BoardState struct {
	Id   string
	Name string
}
type BoardConfig struct {
	BoardStates     []BoardState
	SquadTribes     map[string]string
	TribeChannels   map[string]string
	MandatoryLabels map[string]map[string][]string
	DefaultStateId  string
}

type TicketConfig struct {
	Headers []string
}
