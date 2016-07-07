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

	client := NewLytics(mock.MockApiKey, nil, nil)
	data, err := client.GetSchema("user")

	assert.Equal(t, err, nil)
	assert.Equal(t, data.Name, "user")
	assert.Equal(t, len(data.Columns), 5)
}
