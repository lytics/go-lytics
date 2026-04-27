package lytics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
)

func TestGetProviders(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterProviderMocks()

	client := NewLytics(mock.MockApiKey, nil)
	provider, err := client.GetProviders()
	assert.Equal(t, err, nil)
	assert.True(t, len(provider) > 1)
}

func TestGetProvider(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterProviderMocks()

	client := NewLytics(mock.MockApiKey, nil)
	provider, err := client.GetProvider(mock.MockProviderID)
	assert.Equal(t, err, nil)
	assert.True(t, provider.Id == mock.MockProviderID)
}
