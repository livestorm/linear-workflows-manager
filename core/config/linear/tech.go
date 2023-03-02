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
	ColumnMappings: map[string][]ColumnMapping{
		// These mappings can be configured based on the team of the parent ticket (including the same team as the child).
		// SUP is the key of the parent team in this case.
		"SUP": {
			{ChildColumn: "Backlog", ParentColumn: "Escalated"},
			{ChildColumn: "In Progress", ParentColumn: "In Development"},
			{ChildColumn: "Next release", ParentColumn: "Next Release"},
			{
				ChildColumn:  "Closed",
				ParentColumn: "Closed",
				Comment:      "Hello team, \n The fix of this issue has now been deployed to production.",
			},
		},
	},
	DefaultStateId: environment.Get("DEFAULT_TECH_STATE_ID"), // ID of the TODO State
}
