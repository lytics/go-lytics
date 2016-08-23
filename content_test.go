package lytics

import (
	"github.com/bmizerany/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
	"testing"
)

func TestGetUserContentRecommendation(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterContentMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	recs, err := client.GetUserContentRecommendation("user_id", mock.MockUserID, "", 1, false)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(recs), 1)
	assert.Equal(t, recs[0].Url, "www.testwebsite.com/some/url")
	assert.Equal(t, recs[0].Confidence, 0.8328074038765564)
}

func TestGetSegmentContentRecommendation(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterContentMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	recs, err := client.GetSegmentContentRecommendation(mock.MockSegmentID1, "", 1, false)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(recs), 1)
	assert.Equal(t, recs[0].Url, "www.testwebsite.com/some/url")
	assert.Equal(t, recs[0].Confidence, 0.8328074038765564)
}

func TestGetDocuments(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterContentMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	docs, err := client.GetDocuments([]string{}, 1)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(docs.Urls), 1)
	assert.Equal(t, docs.Total, 5000)
	assert.Equal(t, docs.Urls[0].Url, "www.testwebsite.com/some/url")

	urls := []string{"www.testwebsite.com/some/url", "www.testwebsite.com/path/to/my-article"}
	docs, err = client.GetDocuments(urls, 0)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(docs.Urls), 2)
	assert.Equal(t, docs.Urls[0].Url, "www.testwebsite.com/some/url")
	assert.Equal(t, docs.Urls[1].Url, "www.testwebsite.com/path/to/my-article")
}
