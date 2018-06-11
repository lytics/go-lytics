package lytics

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
	"github.com/stretchr/testify/assert"
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
	assert.NotEqual(t, nil, entity)
	assert.True(t, len(entity.Fields) > 5)
	assert.Equal(t, "sample@sample.com", entity.Fields["email"])

	fields = []string{
		"score_intensity",
	}
	entity, err = client.GetEntity(entitytype, fieldname, fieldval, fields)
	assert.Equal(t, err, nil)
	assert.Equal(t, "83", entity.Fields["score_intensity"])
}
