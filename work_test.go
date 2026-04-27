package lytics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
)

func TestGetWorks(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterWorkMocks()

	client := NewLytics(mock.MockApiKey, nil)
	work, err := client.GetWorks()
	assert.Equal(t, err, nil)
	assert.True(t, len(work) > 1)
}

func TestGetWork(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterWorkMocks()

	client := NewLytics(mock.MockApiKey, nil)
	work, err := client.GetWork(mock.MockWorkID, false)
	assert.Equal(t, err, nil)
	assert.True(t, work.Id == mock.MockWorkID)
}
