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

	filter := &RecommendationFilter{
		Limit:   1,
		Shuffle: false,
		Topics:  []string{"mock", "test"},
		Visited: false,
	}

	client := NewLytics(mock.MockApiKey, nil, nil)
	recs, err := client.GetUserContentRecommendation("user_id", mock.MockUserID, filter)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(recs), 1)
	assert.Equal(t, recs[0].Url, "www.testwebsite.com/some/url")
	assert.Equal(t, recs[0].Confidence, 0.8328074038765564)
}

func TestGetSegmentContentRecommendation(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterContentMocks()

	filter := &RecommendationFilter{
		Limit: 1,
		Rank:  "popular",
	}

	client := NewLytics(mock.MockApiKey, nil, nil)
	recs, err := client.GetSegmentContentRecommendation(mock.MockSegmentID1, filter)

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

func TestGetTopicSummary(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterContentMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	topicSummary, err := client.GetTopicSummary(mock.MockTopicID, 0)

	assert.Equal(t, err, nil)
	assert.Equal(t, topicSummary.Docs.Total, 1)
	assert.Equal(t, topicSummary.Topics.Present, 7890)
	assert.Equal(t, topicSummary.Topics.HighBucket, 351)
}

func TestGetContentTaxonomy(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterContentMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	taxonomy, err := client.GetContentTaxonomy()

	assert.Equal(t, err, nil)
	assert.Equal(t, taxonomy.DocCount, 438)
	assert.Equal(t, taxonomy.Nodes[0].Name, "marketing technology")
	assert.Equal(t, taxonomy.Nodes[1].Count, 1)
}

func TestGetTopicRollups(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterContentMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	rollups, err := client.GetTopicRollups()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(rollups), 2)
	assert.Equal(t, rollups[0].Label, "Marketing")
	assert.Equal(t, len(rollups[0].Topics), 3)
	assert.Equal(t, rollups[0].AcctId, mock.MockAccountID)

	assert.Equal(t, rollups[1].Label, "Data science")
	assert.Equal(t, rollups[1].Topics[0].Label, "Data")
}
