package linear

import (
	"strings"

	"github.com/livestorm/linear-workflows-manager/internal/workflows"
)

func HandleErrorResponse(err error) workflows.WebhookResponse {
	return workflows.WebhookResponse{
		Success: false,
		Error:   strings.TrimSpace(err.Error()),
	}
}
