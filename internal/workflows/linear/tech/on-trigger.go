package tech

import (
	linearprovider "github.com/livestorm/linear-workflows-manager/internal/providers/linear"
	"github.com/livestorm/linear-workflows-manager/internal/workflows"
	"github.com/livestorm/linear-workflows-manager/internal/workflows/linear"
)

var linearProvider = linearprovider.New()

func OnTrigger(ticket *linear.Ticket) (resp workflows.WebhookResponse) {
	if ticket.Action == "create" {
		return onCreate(ticket)
	} else if ticket.Action == "update" {
		return onUpdate(ticket)
	}

	return workflows.WebhookResponse{}
}
