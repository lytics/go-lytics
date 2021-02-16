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
	assert.Equal(t, model.Features[0].Importance, 0.1671249745886582)
	assert.Equal(t, model.Summary.Conf.FalseNegative, 0)
}

func TestGetSegmentMLModels(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterSegmentMLMocks()

	client := NewLytics(mock.MockApiKey, nil)
	models, err := client.GetSegmentMLModels()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(models), 2)

	first := models[0]
	// assert.T(t, ok)

	second := models[1]
	// assert.T(t, ok)

	assert.Equal(t, first.Name, "all::ly_binge_user")
	assert.Equal(t, first.Features[0].Importance, 0.03761499492078325)
	assert.Equal(t, first.Summary.Mse, 0.002072189145635414)

	assert.Equal(t, second.Name, "all::ly_infrequent_user")
	assert.Equal(t, second.Features[0].Type, "numeric")
	assert.Equal(t, second.Summary.Rsq, 0.9389252943819426)
}
