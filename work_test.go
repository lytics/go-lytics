package lytics

import (
	"github.com/bmizerany/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
	"testing"
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
