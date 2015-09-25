package lytics

import (
	"github.com/bmizerany/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
	"testing"
)

func TestGetAuths(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterAuthMocks()

	client := NewLytics(mock.MockApiKey, nil)
	auths, err := client.GetAuths()
	assert.Equal(t, err, nil)
	assert.T(t, len(auths) > 1)
}

func TestGetAuth(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterAuthMocks()

	client := NewLytics(mock.MockApiKey, nil)
	auth, err := client.GetAuth(mock.MockAuthID)
	assert.Equal(t, err, nil)
	assert.T(t, auth.Id == mock.MockAuthID)
}
