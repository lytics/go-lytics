package lytics

import (
	"github.com/bmizerany/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
	"testing"
)

func TestGetCampaign(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterCampaignMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	data, err := client.GetCampaign(mock.MockCampaignID)

	assert.Equal(t, err, nil)
	assert.Equal(t, data.Id, mock.MockCampaignID)
	assert.Equal(t, data.Name, "Mock Campaign")
	assert.Equal(t, data.Status, "published")
}

func TestGetCampaignList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterCampaignMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	data, err := client.GetCampaignList([]string{})

	assert.Equal(t, err, nil)
	assert.Equal(t, len(data), 2)
	assert.Equal(t, data[0].Name, "Mock Campaign")
	assert.Equal(t, data[0].Status, "published")

	assert.Equal(t, data[1].Name, "Another Mock Campaign")
	assert.Equal(t, data[1].Status, "unpublished")

	data, err = client.GetCampaignList([]string{"published"})

	assert.Equal(t, len(data), 1)
}

func TestGetVariation(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterCampaignMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	data, err := client.GetVariation(mock.MockVariationID)

	assert.Equal(t, err, nil)
	assert.Equal(t, data.Id, mock.MockVariationID)
	assert.Equal(t, data.Variation, 0)
	assert.Equal(t, data.CampaignId, mock.MockCampaignID)
}

func TestGetVariationList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterCampaignMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	data, err := client.GetVariationList()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(data), 2)
	assert.Equal(t, data[0].Id, mock.MockVariationID)
	assert.Equal(t, data[0].Variation, 0)
	assert.Equal(t, data[0].CampaignId, mock.MockCampaignID)

	assert.Equal(t, data[1].Id, "b6d237b4f397495eaf633f9")
	assert.Equal(t, data[1].Variation, 1)
	assert.Equal(t, data[1].CampaignId, mock.MockCampaignID)
}
