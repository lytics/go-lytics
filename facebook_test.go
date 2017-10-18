package lytics

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
)

func TestFacebookLookalikeCreateAndStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterFaceookMocks()

	audience := FBLookalikeAudience{}
	audience.Name = "Mark Test Again"
	audience.OriginAudienceId = "6081729845018"
	spec := FBLookalikeAudienceSpec{}
	audience.LookalikeSpec = spec
	audience.LookalikeSpec.Ratio = 0.08
	audience.LookalikeSpec.Country = "US"
	client := NewLytics(mock.MockApiKey, nil, nil)

	resp, err := client.CreateLookalike(mock.MockFacebookAuthID, mock.MockFacebookAccountID, audience)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp, "6082073235818")

	status, err := client.GetLookalikeStatus(mock.MockFacebookAuthID, mock.MockFacebookAudienceID)
	assert.Equal(t, err, nil)
	assert.Equal(t, status.DeliveryStatus.Description, "This audience is ready for use.")
}
