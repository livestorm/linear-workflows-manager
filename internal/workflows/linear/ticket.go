package linear

import (
	"fmt"
	"time"
)

type Ticket struct {
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"createdAt"`
	Data      struct {
		ID                  string        `json:"id"`
		CreatedAt           time.Time     `json:"createdAt"`
		UpdatedAt           time.Time     `json:"updatedAt"`
		Number              int           `json:"number"`
		Title               string        `json:"title"`
		Description         string        `json:"description"`
		Priority            int           `json:"priority"`
		BoardOrder          int           `json:"boardOrder"`
		TeamID              string        `json:"teamId"`
		PreviousIdentifiers []interface{} `json:"previousIdentifiers"`
		CreatorID           string        `json:"creatorId"`
		AssigneeID          string        `json:"assigneeId"`
		StateID             string        `json:"stateId"`
		PriorityLabel       string        `json:"priorityLabel"`
		DescriptionData     string        `json:"descriptionData"`
		SubscriberIds       []string      `json:"subscriberIds"`
		LabelIds            []string      `json:"labelIds"`
		Assignee            struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"assignee"`
		State struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Color string `json:"color"`
			Type  string `json:"type"`
		} `json:"state"`
		Team struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Key  string `json:"key"`
		} `json:"team"`
		Labels []struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Color string `json:"color"`
		} `json:"labels"`
	} `json:"data"`
	UpdatedFrom struct {
		UpdatedAt   time.Time `json:"updatedAt"`
		CompletedAt *string   `json:"completedAt"`
		StateID     string    `json:"stateId"`
	} `json:"updatedFrom"`
	URL            string `json:"url"`
	Type           string `json:"type"`
	OrganizationID string `json:"organizationId"`
}

func (t *Ticket) TType() string {
	return t.GetLabel("type")
}

func (t *Ticket) Squad() string {
	return t.GetLabel("squad")
}

func (t *Ticket) Identifier() string {
	return fmt.Sprintf("%s-%d", t.Data.Team.Key, t.Data.Number)
}
