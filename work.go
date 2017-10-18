package lytics

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

const (
	workEndpoint       = "work/:id" // state
	workListEndpoint   = "work"
	workCreateEndpoint = "work"
)

type Work struct {
	Id               string                              `json:"id"`
	Name             string                              `json:"name,omitempty"`
	Description      string                              `json:"description,omitempty"`
	Aid              int                                 `json:"aid"`
	AccountId        string                              `json:"account_id"`
	UserId           string                              `json:"user_id"`
	WorkStateVersion int                                 `json:"work_state_version"`
	Config           json.RawMessage                     `json:"config,omitempty"`
	WorkflowId       string                              `json:"workflow_id"`
	Workflow         string                              `json:"workflow"`
	Tag              string                              `json:"tag"`
	Updated          time.Time                           `json:"updated"`
	Created          time.Time                           `json:"created"`
	AuthIds          []string                            `json:"auth_ids,omitempty"`
	Hidden           bool                                `json:"hidden"`
	RuntimeOverride  string                              `json:"runtime"`
	VerboseLogging   bool                                `json:"verbose_logging"`
	StatusCode       string                              `json:"statuscode"`
	Progress         map[string][]map[string]interface{} `json:"progress"`
}

// GetWork returns a single work
// https://www.getlytics.com/developers/rest-api#work-get
func (l *Client) GetWork(id string, state bool) (Work, error) {
	res := ApiResp{}
	data := Work{}

	params := url.Values{}
	params.Add("state", strconv.FormatBool(state))

	// make the request
	err := l.Get(parseLyticsURL(workEndpoint, map[string]string{"id": id}), params, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// GetWorks returns all works
// https://www.getlytics.com/developers/rest-api#work
func (l *Client) GetWorks() ([]Work, error) {
	res := ApiResp{}
	data := []Work{}

	// make the request
	err := l.Get(workListEndpoint, nil, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// CreateWork creates a single work from config
func (l *Client) CreateWork(work Work) (Work, error) {
	var (
		err error
	)

	res := ApiResp{}
	data := Work{}

	err = l.Post(workCreateEndpoint, nil, work, &res, &data)
	if err != nil {
		return work, err
	}

	return data, nil
}
