package lytics

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
)

func TestGetUsers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterUserMocks()
	mock.RegisterAccountMocks()

	client := NewLytics(mock.MockApiKey, nil)
	users, err := client.GetUsers()
	assert.Equal(t, err, nil)
	assert.T(t, len(users) > 1)
}

func TestGetUser(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterUserMocks()
	mock.RegisterAccountMocks()

	client := NewLytics(mock.MockApiKey, nil)
	user, err := client.GetUser(mock.MockUserID)
	assert.Equal(t, err, nil)
	assert.T(t, user.Id == mock.MockUserID)
}
