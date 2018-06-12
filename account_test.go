package lytics

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
)

func TestGetAccounts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterAccountMocks()

	client := NewLytics(mock.MockApiKey, nil)
	accts, err := client.GetAccounts()
	assert.Equal(t, err, nil)
	assert.T(t, len(accts) > 1)
}

func TestGetAccount(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterAccountMocks()

	client := NewLytics(mock.MockApiKey, nil)
	acct, err := client.GetAccount(mock.MockAccountID)
	assert.Equal(t, err, nil)
	assert.T(t, acct.Id == mock.MockAccountID)
}
