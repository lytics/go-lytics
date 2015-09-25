package lytics

import (
	"github.com/bmizerany/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
	"testing"
)

func TestGetSchema(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterCatalogMocks()

	client := NewLytics(mock.MockApiKey, nil)
	data, err := client.GetSchema("user", 2)

	assert.Equal(t, err, nil)
	assert.Equal(t, data.Name, "user")
	assert.Equal(t, len(data.Columns), 2)

	client = NewLytics(mock.MockApiKey, nil)
	data, err = client.GetSchema("user", -1)

	assert.Equal(t, err, nil)
	assert.Equal(t, data.Name, "user")
	assert.Equal(t, len(data.Columns), 5)
}
