package lytics

import (
	"github.com/bmizerany/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
	"testing"
)

func TestGetEntity(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterEntityMocks()

	client := NewLytics(mock.MockApiKey, nil)

	fields := []string{}
	entitytype := "user"
	fieldname := "email"
	fieldval := "sample@sample.com"

	entity, err := client.GetEntity(entitytype, fieldname, fieldval, fields)
	assert.Equal(t, err, nil)
	assert.T(t, len(entity) > 5)
	assert.T(t, entity["email"].(string) == "sample@sample.com")

	fields = []string{
		"score_intensity",
	}
	entity, err = client.GetEntity(entitytype, fieldname, fieldval, fields)
	assert.Equal(t, err, nil)
	assert.Equal(t, entity["score_intensity"].(string), "83")
}
