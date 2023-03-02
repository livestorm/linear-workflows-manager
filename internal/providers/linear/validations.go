package linear

import (
	linearconfig "github.com/livestorm/linear-workflows-manager/core/config/linear"
	"github.com/livestorm/linear-workflows-manager/internal/workflows/linear"
)

func GetBoardPositionById(sl []linearconfig.BoardState, id string) int {
	for idx, state := range sl {
		if state.Id == id {
			return idx
		}
	}
	return -1
}

func GetBoardPositionByName(sl []linearconfig.BoardState, name string) int {
	for idx, state := range sl {
		if state.Name == name {
			return idx
		}
	}
	return -1
}

func (l *LinearProvider) CompareBoardPosition(teamKey string, oldStateId string, newStateId string) int {
	config := linearconfig.GetBoardConfig(teamKey)
	oldStatePosition := GetBoardPositionById(config.BoardStates, oldStateId)
	newStatePosition := GetBoardPositionById(config.BoardStates, newStateId)
	if oldStatePosition == -1 || newStatePosition == -1 {
		return -1
	}
	return newStatePosition - oldStatePosition
}

func (l *LinearProvider) CheckMandatoryLabels(ticket *linear.Ticket) ([]string, []string) {
	labelsRequiredNow, missingLabelsRequiredNow := []string{}, []string{}
	labelsRequiredLater, missingLabelsRequiredLater := []string{}, []string{}

	// Load Board Configuration
	config := linearconfig.GetBoardConfig(ticket.Data.Team.Key)

	// Get current status position
	currentStatePosition := GetBoardPositionByName(config.BoardStates, ticket.Data.State.Name)
	if currentStatePosition == -1 {
		return missingLabelsRequiredNow, missingLabelsRequiredLater
	}

	for state, typeLabels := range config.MandatoryLabels {
		statePosition := GetBoardPositionByName(config.BoardStates, state)
		// Get all relevant labels for current state
		stateLabels := typeLabels["global"]
		stateLabels = append(stateLabels, typeLabels[ticket.TType()]...)

		// If current state is after the state, the labels are mandatory else simply warn
		if currentStatePosition >= statePosition {
			labelsRequiredNow = append(labelsRequiredNow, stateLabels...)
		} else {
			labelsRequiredLater = append(labelsRequiredLater, stateLabels...)
		}
	}

	// Iterate through mandatory labels ensuring their existence
	for _, labelRequiredNow := range labelsRequiredNow {
		if ticket.GetLabel(labelRequiredNow) == "" {
			missingLabelsRequiredNow = append(missingLabelsRequiredNow, labelRequiredNow)
		}
	}

	// Iterate through Warning labels
	for _, labelRequiredLater := range labelsRequiredLater {
		if ticket.GetLabel(labelRequiredLater) == "" {
			missingLabelsRequiredLater = append(missingLabelsRequiredLater, labelRequiredLater)
		}
	}

	return missingLabelsRequiredNow, missingLabelsRequiredLater
}

func (l *LinearProvider) HandleColumnMappings(ticket *linear.Ticket) error {
	if ticket.Data.ParentID == "" {
		return nil
	}

	config := linearconfig.GetBoardConfig(ticket.Data.Team.Key)
	parentTicket, err := l.GetIssueById(ticket.Data.ParentID)
	if err != nil {
		return err
	}

	// Check if mapping exists
	colMappings, exists := config.ColumnMappings[parentTicket.Team.Key]
	if !exists {
		return nil
	}

	// Get position of current ticket and see if it matches a mapping
	currentStatePosition := GetBoardPositionByName(config.BoardStates, ticket.Data.State.Name)
	if currentStatePosition == -1 {
		return nil
	}

	// Change State
	newStateId, comment := parentTicket.State.ID, ""

	// Iterate through columns in the map and match with current position
	for _, mapping := range colMappings {
		parentState := l.GetStateByName(parentTicket.Team.Key, mapping.ParentColumn)
		if parentState.Id == "" {
			continue
		}

		currentMappingPosition := GetBoardPositionByName(config.BoardStates, mapping.ChildColumn)
		if currentStatePosition >= currentMappingPosition {
			newStateId = parentState.Id
			comment = mapping.Comment
		}
	}

	if newStateId != parentTicket.State.ID {
		// create comment if required for state change
		if len(comment) > 0 {
			err := l.AddComment(parentTicket.ID, comment)
			if err != nil {
				return err
			}
		}

		// change the state of the ticket
		err := l.ChangeState(parentTicket.ID, newStateId)
		if err != nil {
			return err
		}
	}

	return nil
}
