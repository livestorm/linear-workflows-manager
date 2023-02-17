package tech

import (
	"fmt"

	linearconfig "github.com/livestorm/linear-workflows-manager/core/config/linear"
	"github.com/livestorm/linear-workflows-manager/internal/workflows"
	"github.com/livestorm/linear-workflows-manager/internal/workflows/linear"
)

func checkMandatoryLabelsOnCreate(ticket *linear.Ticket, boardConfig linearconfig.BoardConfig) error {
	missingLabelsRequiredNow, missingLabelsRequiredLater := linearProvider.CheckMandatoryLabels(ticket)
	allMissingLabels := append(missingLabelsRequiredNow, missingLabelsRequiredLater...)

	// No missing labels, so just return.
	if len(allMissingLabels) <= 0 {
		return nil
	}

	commentMessage := fmt.Sprintf(
		"This ticket requires the following labels: `%v`. These may block your ticket's progress in the future.",
		allMissingLabels,
	)

	if len(missingLabelsRequiredNow) > 0 {
		commentMessage += fmt.Sprintf("\n\nCritical: Due to the these missing labels: `%v`, your status is being moved back to the default status: To Spec.", missingLabelsRequiredNow)
		err := linearProvider.ChangeState(ticket.Identifier(), boardConfig.DefaultStateId)
		if err != nil {
			return err
		}
	}

	return linearProvider.AddComment(ticket.Identifier(), commentMessage)
}

func onCreate(ticket *linear.Ticket) (resp workflows.WebhookResponse) {
	// Load board configuration
	boardConfig := linearconfig.GetBoardConfig(ticket.Data.Team.Key)
	err := checkMandatoryLabelsOnCreate(ticket, boardConfig)
	if err != nil {
		return linear.HandleErrorResponse(err)
	}
	resp.Success = true
	return resp
}
