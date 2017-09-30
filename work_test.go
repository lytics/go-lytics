package lytics

import (
	"strings"
	"testing"

	"encoding/json"

	"github.com/bmizerany/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
)

func TestGetWorks(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterWorkMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	work, err := client.GetWorks()
	assert.Equal(t, err, nil)
	assert.T(t, len(work) > 1)
}

func TestGetWork(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterWorkMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	work, err := client.GetWork(mock.MockWorkID, false)
	assert.Equal(t, err, nil)
	assert.T(t, work.Id == mock.MockWorkID)
}

func TestCreateWork(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterWorkMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)

	// test work success
	// create sample config
	config := make(map[string]interface{})
	config["ads_account_id"] = "10150970483556602"
	config["identifier_field"] = "email"
	config["segment_id"] = "9090c9c7f540418a9ee8eb4c62935c56"
	config["version"] = "1"
	config["withbackfill"] = true
	configJSON, _ := json.Marshal(config)

	// build the work model
	myWork := Work{
		Name:       "My Test Work",
		Config:     configJSON,
		WorkflowId: "9151ddcc8961d4bf7ec1cba02899c123",
		AuthIds: []string{
			"84a151a3c32aa2ee9c5479076b20c411",
		},
	}

	work, err := client.CreateWork(myWork)
	assert.Equal(t, err, nil)
	assert.Equal(t, work.Id, "18594ed78384fc48056c5650d6ffe56f")
	assert.Equal(t, work.Workflow, "facebook_audience_populate")

	// test work fail - duplicate
	// create sample config
	config = make(map[string]interface{})
	config["ads_account_id"] = "10150970483556602"
	config["identifier_field"] = "email"
	config["segment_id"] = "iamaduplicatesegmentandshouldfail"
	config["version"] = "1"
	config["withbackfill"] = true
	configJSON, _ = json.Marshal(config)

	// build the work model
	myWork = Work{
		Name:       "My Test Work",
		Config:     configJSON,
		WorkflowId: "9151ddcc8961d4bf7ec1cba02899c123",
		AuthIds: []string{
			"84a151a3c32aa2ee9c5479076b20c411",
		},
	}

	work, err = client.CreateWork(myWork)
	assert.NotEqual(t, err, nil)
	assert.T(t, strings.Contains(err.Error(), "Duplicate work config"))

	// test work fail - unauthorized
	client = NewLytics("notarealkey", nil, nil)

	// build the work model
	myWork = Work{
		Name:       "My Test Work",
		Config:     configJSON,
		WorkflowId: "9151ddcc8961d4bf7ec1cba02899c123",
	}

	work, err = client.CreateWork(myWork)
	assert.NotEqual(t, err, nil)
	assert.T(t, strings.Contains(err.Error(), "Not authorized"))

}
