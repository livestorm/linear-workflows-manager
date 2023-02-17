package linear

import (
	"net/http"

	linearconfig "github.com/livestorm/linear-workflows-manager/core/config/linear"
)

type LinearProvider struct {
	baseUrl string
	token   string
	client  http.Client
}

type linearUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type LinearState struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type getUserResponse struct {
	Data struct {
		User linearUser `json:"user"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type getStatesResponse struct {
	Data struct {
		WorkflowStates struct {
			Nodes []linearconfig.BoardState `json:"nodes"`
		} `json:"workflowStates"`
		User linearUser `json:"user"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type mutationResponse struct {
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type linearRequest struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}
