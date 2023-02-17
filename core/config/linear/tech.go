package linear

import "github.com/livestorm/linear-workflows-manager/core/environment"

var techBoardConfig = BoardConfig{
	BoardStates: []BoardState{
		{"a3aae219-1110-4585-a056-56596088f742", "Canceled"},
		{"17f4a4ae-629b-431a-8178-55b0223d04da", "Backlog"},
		{"ae3ac87f-e113-4fc2-a4bf-256c554e98e6", "Todo"},
		{"8d70eb64-1143-4ee4-afee-0a1b103bfb7d", "In Progress"},
		{"0c205e69-ca02-4e69-b4f8-c8ea95e15088", "Ready for deploy"},
		{"4c37d2db-d251-4719-94f7-31336f52cdf0", "Next release"},
		{"dca387ad-d568-48b1-8b1c-43a917936c2d", "Done"},
		{"568507e9-2013-41fd-a64f-5bb7cdadec06", "Closed"},
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
