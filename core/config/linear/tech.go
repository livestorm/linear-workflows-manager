package linear

import "github.com/livestorm/linear-workflows-manager/core/environment"

var techBoardConfig = BoardConfig{
	BoardStates: []BoardState{
		// Add board states here getting it from the Linear API manually
	},
	MandatoryLabels: map[string]map[string][]string{
		"In Progress": {
			"global": {"type", "scope", "squad", "chapter"},
		},
		"Ready for deploy": {
			"bug":        {"bug-reason"},
			"regression": {"bug-reason"},
		},
	},
	DefaultStateId: environment.Get("DEFAULT_TECH_STATE_ID"), // ID of the TODO State
}
