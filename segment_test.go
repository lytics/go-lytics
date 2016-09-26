package lytics

import (
	"fmt"
	"github.com/bmizerany/assert"
	"github.com/jarcoal/httpmock"
	"github.com/lytics/go-lytics/mock"
	"testing"
)

func TestGetSegments(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterSegmentMocks()

	// user segment
	client := NewLytics(mock.MockApiKey, nil, nil)
	segs, err := client.GetSegments("user")
	assert.Equal(t, err, nil)
	assert.T(t, len(segs) > 1)
	assert.Equal(t, segs[0].Table, "user")

	// content segment
	segs, err = client.GetSegments("content")
	assert.Equal(t, err, nil)
	assert.T(t, len(segs) > 1)
	assert.Equal(t, segs[0].Table, "content")
}

func TestGetSegment(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterSegmentMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	seg, err := client.GetSegment(mock.MockSegmentID1)
	assert.Equal(t, err, nil)
	assert.T(t, seg.Id == mock.MockSegmentID1)
}

func TestGetSegmentSize(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterSegmentMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	seg, err := client.GetSegmentSize(mock.MockSegmentID1)

	assert.Equal(t, err, nil)
	assert.T(t, seg.Id == mock.MockSegmentID1)
}

func TestGetSegmentSizes(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterSegmentMocks()

	var segments []string

	client := NewLytics(mock.MockApiKey, nil, nil)

	segments = []string{
		mock.MockSegmentID1,
	}

	seg, err := client.GetSegmentSizes(segments)
	assert.Equal(t, err, nil)
	assert.T(t, seg[0].Id == segments[0])

	segments = []string{
		mock.MockSegmentID1,
		mock.MockSegmentID2,
	}

	// params
	seg, err = client.GetSegmentSizes(segments)
	assert.Equal(t, err, nil)
	assert.T(t, len(seg) == 2)
}

func TestGetSegmentAttribution(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterSegmentMocks()

	segments := []string{
		mock.MockSegmentID1,
		mock.MockSegmentID2,
	}

	client := NewLytics(mock.MockApiKey, nil, nil)
	attr, err := client.GetSegmentAttribution(segments)

	assert.Equal(t, err, nil)
	assert.T(t, len(attr[0].Metrics) == 5)
	assert.T(t, len(attr[1].Metrics) == 5)
}

func TestGetSegmentCollection(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterSegmentMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	collection, err := client.GetSegmentCollection(mock.MockSegmentCollection)

	assert.Equal(t, err, nil)
	assert.Equal(t, collection.Slug, "goal_funnel")
	assert.T(t, len(collection.Collection) == 2)
	assert.Equal(t, collection.Collection[0].Id, mock.MockSegmentID1)
	assert.Equal(t, collection.Collection[1].Id, mock.MockSegmentID2)
}

func TestGetSegmentCollectionList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterSegmentMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)
	collections, err := client.GetSegmentCollectionList()

	assert.Equal(t, err, nil)
	assert.T(t, len(collections) == 2)
	assert.Equal(t, collections[0].Slug, "goal_funnel")
	assert.T(t, len(collections[0].Collection) == 2)
	assert.Equal(t, collections[1].Slug, "mobile")
	assert.T(t, len(collections[1].Collection) == 3)
}

func TestSegmentPager(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterSegmentMocks()

	var (
		completed     bool
		countCalls    int
		countEntities int
	)

	client := NewLytics(mock.MockApiKey, nil, nil)

	// create the segment scanner
	err := client.CreateScanner()
	assert.Equal(t, err, nil)

	// start the paging routine
	err = client.PageMembers(mock.MockSegmentID1)
	assert.Equal(t, err, nil)

	// handle processing the entities
PagingComplete:
	for {
		select {
		case entities := <-client.Scan.Loader:
			countCalls++

			for _, v := range entities {
				countEntities++
				assert.Equal(t, v["email"], fmt.Sprintf("email%d@email.com", countEntities))
			}

		case shutdown := <-client.Scan.Shutdown:
			if shutdown {
				completed = true
				break PagingComplete
			}
		}
	}
	assert.Equal(t, countCalls, 3)
	assert.Equal(t, completed, true)
	fmt.Printf("*** COMPLETED SCAN: Loaded %d batches and %d total entities", len(client.Scan.Batches), client.Scan.Total)
}

func TestCreateSegment(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterSegmentMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)

	ql := "FILTER AND (created >= \"now-30d\", aspects = \"articles\") FROM content"

	seg, err := client.CreateSegment("Recent Articles", ql, "recent_articles")
	assert.Equal(t, err, nil)
	assert.T(t, seg.Id != "")
	assert.Equal(t, seg.FilterQL, ql)

	seg, err = client.CreateSegment("Invalid Segment", "", "invalid")
	assert.NotEqual(t, err, nil)
}

func TestValidateSegment(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mock.RegisterSegmentMocks()

	client := NewLytics(mock.MockApiKey, nil, nil)

	valid, err := client.ValidateSegment("FILTER AND (created >= \"now-30d\", aspects = \"articles\") FROM content")
	assert.Equal(t, err, nil)
	assert.T(t, valid)

	valid, err = client.ValidateSegment("")
	assert.NotEqual(t, err, nil)
	assert.T(t, !valid)
}
