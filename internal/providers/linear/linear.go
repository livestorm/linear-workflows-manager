package linear

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	linearconfig "github.com/livestorm/linear-workflows-manager/core/config/linear"
	"github.com/livestorm/linear-workflows-manager/core/environment"
)

func New() *LinearProvider {
	return &LinearProvider{
		baseUrl: environment.Get("LINEAR_API_BASE_URL") + "/graphql",
		token:   environment.Get("LINEAR_API_TOKEN"),
		client: http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (l *LinearProvider) buildRequest(payload interface{}) (*http.Request, error) {
	var reader io.Reader
	if payload != nil {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(payload); err != nil {
			return nil, err
		}
		reader = &buf
	}

	request, err := http.NewRequest("POST", l.baseUrl, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", l.token)
	request.Header.Add("content-type", "application/json; charset=utf-8")
	return request, nil
}

func (l *LinearProvider) GetUserById(id string) (user *linearUser, err error) {
	// Build payload
	payload := &linearRequest{}
	payload.Query = getUserQuery

	payload.Variables = struct {
		Identifier string `json:"identifier"`
	}{
		Identifier: id,
	}

	// Create request and add headers
	req, err := l.buildRequest(payload)
	if err != nil {
		fmt.Printf("Failed to create HTTP request: %s \n", err.Error())
		return nil, err
	}

	// Perform the request
	resp, err := l.client.Do(req)
	if err != nil {
		fmt.Printf("Failed to run HTTP request: %s \n", err.Error())
		return nil, err
	}

	// Parse the response received
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read body: %s \n", err.Error())
		return nil, err
	}

	response := &getUserResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		fmt.Printf("Failed to parse response body: %s \n", err.Error())
		return nil, err
	}

	// Check if successful and return response
	if len(response.Errors) > 0 {
		fmt.Printf("Linear Query failed: %s \n", response.Errors[0].Message)
		return nil, errors.New(response.Errors[0].Message)
	}

	user = &response.Data.User
	return user, nil
}

func (l *LinearProvider) GetStatesByTeam(name string) (states *[]linearconfig.BoardState, err error) {
	// Build payload
	payload := &linearRequest{}
	payload.Query = getStatesQuery

	payload.Variables = struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	// Create request and add headers
	req, err := l.buildRequest(payload)
	if err != nil {
		fmt.Printf("Failed to create HTTP request: %s \n", err.Error())
		return nil, err
	}

	// Perform the request
	resp, err := l.client.Do(req)
	if err != nil {
		fmt.Printf("Failed to run HTTP request: %s \n", err.Error())
		return nil, err
	}

	// Parse the response received
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read body: %s \n", err.Error())
		return nil, err
	}

	response := &getStatesResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		fmt.Printf("Failed to parse response body: %s \n", err.Error())
		return nil, err
	}

	// Check if successful and return response
	if len(response.Errors) > 0 {
		fmt.Printf("Linear Query failed: %s \n", response.Errors[0].Message)
		return nil, errors.New(response.Errors[0].Message)
	}

	states = &response.Data.WorkflowStates.Nodes
	return states, nil
}

func (l *LinearProvider) executeMutation(payload interface{}) error {
	// Create request and add headers
	req, err := l.buildRequest(payload)
	if err != nil {
		fmt.Printf("Failed to create HTTP request: %s \n", err.Error())
		return err
	}

	// Perform the request
	resp, err := l.client.Do(req)
	if err != nil {
		fmt.Printf("Failed to run HTTP request: %s \n", err.Error())
		return err
	}

	// Parse the response received
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read body: %s \n", err.Error())
		return err
	}

	log.Print(string(body))
	response := &mutationResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		fmt.Printf("Failed to parse response body: %s \n", err.Error())
		return err
	}

	if len(response.Errors) > 0 {
		return errors.New(response.Errors[0].Message)
	}

	return nil
}

func (l *LinearProvider) AddComment(ticketId string, comment string) error {
	payload := &linearRequest{}
	payload.Query = addCommentMutation
	payload.Variables = struct {
		Identifier string `json:"identifier"`
		Comment    string `json:"comment"`
	}{
		Identifier: ticketId,
		Comment:    comment,
	}

	// Run request and parse response
	err := l.executeMutation(payload)
	return err
}

func (l *LinearProvider) ChangeState(ticketId string, stateId string) error {
	payload := &linearRequest{}
	payload.Query = changeStateMutation

	payload.Variables = struct {
		Identifier string `json:"identifier"`
		StateId    string `json:"stateId"`
	}{
		Identifier: ticketId,
		StateId:    stateId,
	}

	log.Print(stateId)

	// Run request and parse response
	err := l.executeMutation(payload)
	return err
}
