package workflows

type WebhookResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
