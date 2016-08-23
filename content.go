package lytics

import (
	"strconv"
	"strings"
	"time"
)

const (
	userRecommendEndpoint    = "content/recommend/user/:fieldName/:fieldVal"
	segmentRecommendEndpoint = "content/recommend/segment/:id"
	documentsEndpoint        = "content/doc"
)

type Document struct {
	Url             string             `json:"url"`
	Title           string             `json:"title"`
	Description     string             `json:"description"`
	Topics          []string           `json:"topics"`
	TopicRelevances map[string]float64 `json:"topic_relevances"`
	PrimaryImage    string             `json:"primary_image"`
	Author          string             `json:"author"`
	Created         time.Time          `json:"created"`
	Hash            string             `json:"id"`
	Sitename        string             `json:"sitename,omitempty"`
	Stream          string             `json:"stream"`
	Path            []string           `json:"path,omitempty"`
	Aspects         []string           `json:"aspects,omitempty"`
	Language        string             `json:"language,omitempty"`
	Updated         time.Time          `json:"updated,omitempty"`
	Fetched         time.Time          `json:"fetched,omitempty"`
	Meta            []string           `json:"meta,omitempty"`
	Id              string             `json:"hashedurl,omitempty"`
}

type Recommendation struct {
	*Document
	Confidence float64 `json:"confidence"`
	Visited    bool    `json:"visited"`
	VisitRate  float64 `json:"visitrate,omitempty"`
}

type Documents struct {
	Urls  []Document `json: "urls"`
	Total int        `json: "total"`
}

// GetUserContentRecommendation returns a list of documents
// to recommend the user based on their content affinities
// Example QL string: FILTER AND (url LIKE "www.example.com/*") FROM content
func (l *Client) GetUserContentRecommendation(fieldName, fieldVal, ql string, limit int, shuffle bool) ([]Recommendation, error) {
	res := ApiResp{}
	data := []Recommendation{}
	params := map[string]string{}

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	if shuffle {
		params["shuffle"] = "true"
	}

	if ql != "" {
		params["ql"] = ql
	}

	// make the request
	err := l.Get(parseLyticsURL(userRecommendEndpoint, map[string]string{"fieldName": fieldName, "fieldVal": fieldVal}), params, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// GetSegmentContentRecommendation returns a list of documents
// to recommend to users in a segment
func (l *Client) GetSegmentContentRecommendation(segId string, ql string, limit int, shuffle bool) ([]Recommendation, error) {
	res := ApiResp{}
	data := []Recommendation{}
	params := map[string]string{}

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	if shuffle {
		params["shuffle"] = "true"
	}

	if ql != "" {
		params["ql"] = ql
	}

	// make the request
	err := l.Get(parseLyticsURL(segmentRecommendEndpoint, map[string]string{"id": segId}), params, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// GetDocuments returns a summary for document(s).
// If no url is specified, then return the 10 most recent urls.
// https://www.getlytics.com/developers/rest-api#content-document
func (l *Client) GetDocuments(urls []string, limit int) (Documents, error) {
	res := ApiResp{}
	data := Documents{}
	params := map[string]string{}

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	if len(urls) > 0 {
		params["urls"] = strings.Join(urls, ",")
	}

	err := l.Get(documentsEndpoint, params, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
