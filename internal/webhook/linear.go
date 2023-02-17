package webhook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/livestorm/linear-workflows-manager/core/environment"
	"github.com/livestorm/linear-workflows-manager/internal/workflows"
	"github.com/livestorm/linear-workflows-manager/internal/workflows/linear"
	"github.com/livestorm/linear-workflows-manager/internal/workflows/linear/tech"
)

func isValidLinearSource(req *http.Request) bool {
	// Check IP
	authorizedIps := environment.Get("LINEAR_ALLOWED_IPS")
	sourceIP := req.Header.Get("X-Forwarded-For")
	if sourceIP == "" || !strings.Contains(authorizedIps, sourceIP) {
		return false
	}

	// Check Headers
	if req.Header.Get("user-agent") != "Linear-Webhook" ||
		req.Header.Get("linear-event") != "Issue" ||
		req.Header.Get("linear-delivery") == "" {
		return false
	}

	return true
}

func processIncomingWebhook(req *http.Request) (*linear.Ticket, error) {
	payload := &linear.Ticket{}

	if !isValidLinearSource(req) {
		return nil, errors.New("Invalid source. Only requests from Linear are allowed.")
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Could not read body: %s \n", err.Error())
		return nil, err
	}

	if err := json.Unmarshal(body, payload); err != nil {
		return nil, err
	}

	return payload, nil
}

func (a *apiRouter) handleLinearTechWebhook(rw http.ResponseWriter, req *http.Request) {
	resp := workflows.WebhookResponse{}
	payload, err := processIncomingWebhook(req)
	if err != nil {
		resp.Error = err.Error()
		a.WriteJson(rw, http.StatusBadRequest, resp)
		return
	}

	resp = tech.OnTrigger(payload)
	if resp.Success == false {
		log.Printf(resp.Error)
	}
	a.WriteJson(rw, http.StatusOK, resp)
}

func (a *apiRouter) linearRoutes() []Route {
	// We defined a route for each of our teams on Linear
	return []Route{
		{
			Method:      http.MethodPost,
			Path:        "/webhooks/linear/tech",
			HandlerFunc: a.handleLinearTechWebhook,
		},
	}
}
