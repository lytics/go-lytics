package lytics

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
)

func TestGetSegmentMLModel(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterSegmentMLMocks()

	client := NewLytics(mock.MockApiKey, nil)
	model, err := client.GetSegmentMLModel(mock.MockSegmentMLID)

	assert.Equal(t, err, nil)
	assert.Equal(t, model.Name, mock.MockSegmentMLID)
	assert.Equal(t, model.Importance["lytics_score_frequency"], 0.5850666666666667)
	assert.Equal(t, model.Summary.Conf.FalseNegative, 3)
}

func TestGetSegmentMLModels(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterSegmentMLMocks()

	client := NewLytics(mock.MockApiKey, nil)
	models, err := client.GetSegmentMLModels()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(models), 2)

	first, ok := models[mock.MockSegmentMLID]
	assert.T(t, ok)

	second, ok := models["goal_anonymous::goal_known"]
	assert.T(t, ok)

	assert.Equal(t, first.Importance["lytics_score_recency"], 0.10316666666666663)
	assert.Equal(t, first.Summary.Mse, 0.27024152611111113)

	assert.Equal(t, second.Name, "goal_anonymous::goal_known")
	assert.Equal(t, second.FieldTypes["goal_known"], "numeric")
}
