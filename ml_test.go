package lytics

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
)

func TestGetMLModel(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterMLMocks()

	client := NewLytics(mock.MockApiKey, nil)
	model, err := client.GetMLModel(mock.MockMLID)

	assert.Equal(t, err, nil)
	assert.Equal(t, model.ID, mock.MockMLID)
	assert.Equal(t, model.Config.AutoTune, false)
}

func TestGetMLModels(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterMLMocks()

	client := NewLytics(mock.MockApiKey, nil)
	models, err := client.GetMLModels()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(models), 2)

	first := models[0]

	assert.Equal(t, first.IsActive, false)
	assert.Equal(t, first.Name, "all::ly_binge_user")

	second := models[1]

	assert.Equal(t, second.Name, "all::ly_infrequent_user")
	assert.Equal(t, second.IsHealthy, true)
}
