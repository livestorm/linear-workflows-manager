package tech

import (
	"fmt"

	linearconfig "github.com/livestorm/linear-workflows-manager/core/config/linear"
	"github.com/livestorm/linear-workflows-manager/internal/workflows"
	"github.com/livestorm/linear-workflows-manager/internal/workflows/linear"
)

func checkMandatoryLabelsOnUpdate(ticket *linear.Ticket, boardConfig linearconfig.BoardConfig) error {
	missingLabelsRequiredNow, _ := linearProvider.CheckMandatoryLabels(ticket)
	if len(missingLabelsRequiredNow) > 0 {
		commentMessage := fmt.Sprintf(
			"*Status update failed:* Ticket does not contain the required labels: `%v` to be in %s",
			missingLabelsRequiredNow,
			ticket.Data.State.Name,
		)
		err := linearProvider.AddComment(ticket.Identifier(), commentMessage)
		if err != nil {
			return err
		}

		previousState := ticket.UpdatedFrom.StateID
		positionDiff := linearProvider.CompareBoardPosition(ticket.Data.Team.Key, ticket.Data.State.ID, previousState)
		// We should only allow the position moving backwards and not forward to avoid infinite loop
		if positionDiff >= 0 {
			previousState = boardConfig.DefaultStateId
		}
		return linearProvider.ChangeState(ticket.Identifier(), previousState)
	}
	return nil
}

func handleStateChange(ticket *linear.Ticket) error {
	// Load board configuration
	boardConfig := linearconfig.GetBoardConfig(ticket.Data.Team.Key)

	// Check Mandatory Labels
	err := checkMandatoryLabelsOnUpdate(ticket, boardConfig)
	if err != nil {
		return err
	}
	return nil
}

func onUpdate(ticket *linear.Ticket) (resp workflows.WebhookResponse) {
	if ticket.UpdatedFrom.StateID != "" {
		err := handleStateChange(ticket)
		if err != nil {
			return linear.HandleErrorResponse(err)
		}
	}
	resp.Success = true
	return resp
}
