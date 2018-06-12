package lytics

import (
	"testing"

	"github.com/bmizerany/assert"
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
	assert.T(t, len(provider) > 1)
}

func TestGetProvider(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterProviderMocks()

	client := NewLytics(mock.MockApiKey, nil)
	provider, err := client.GetProvider(mock.MockProviderID)
	assert.Equal(t, err, nil)
	assert.T(t, provider.Id == mock.MockProviderID)
}
