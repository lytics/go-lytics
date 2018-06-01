package lytics

import (
	"encoding/json"
	"net/url"
	"time"
)

const (
	queryEndpoint         = "query/:id"
	queryListEndpoint     = "query"
	queryTestEndpoint     = "query/_test"
	queryValidateEndpoint = "query/_validate"
)

type (
	// Query represents an LQL Statement structure
	Query struct {
		Id      string                `json:"id"`
		Created time.Time             `json:"created"`
		Updated time.Time             `json:"updated"`
		Alias   string                `json:"alias"`
		Table   string                `json:"table"`
		From    string                `json:"from"`
		Text    string                `json:"text"`
		Fields  map[string]QueryField `json:"fields"`
	}
	// A field in a query
	// - very similar to catalog, query fields create catalog fields
	QueryField struct {
		As         string   `json:"as"`
		IsBy       bool     `json:"is_by"`
		Type       string   `json:"type"`
		ShortDesc  string   `json:"shortdesc"`
		LongDesc   string   `json:"longdesc"`
		Identities []string `json:"identities"`
		Froms      []string `json:"froms"`
	}
)

// GetQueries returns a list of all queries associated with this account
// https://www.getlytics.com/developers/rest-api#query
func (l *Client) GetQueries() ([]Query, error) {
	res := ApiResp{}
	data := []Query{}

	// make the request
	err := l.Get(queryListEndpoint, nil, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// GetQueries returns a list of all queries associated with this account
// https://www.getlytics.com/developers/rest-api#query
func (l *Client) GetQuery(alias string) (Query, error) {
	res := ApiResp{}
	data := Query{}

	// make the request
	err := l.Get(parseLyticsURL(queryEndpoint, map[string]string{"id": alias}), nil, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// GetQueryTest returns the evaluated entity from given query
// https://www.getlytics.com/developers/rest-api#query
func (l *Client) GetQueryTest(qs url.Values, query string) (Entity, error) {
	res := ApiResp{}
	data := Entity{}

	// make the request
	err := l.Post(queryTestEndpoint, qs, query, &res, &data)

	if err != nil {
		return data, err
	}

	return data, nil
}

// PostQueryValidate returns the query and how it is interpreted
// https://www.getlytics.com/developers/rest-api#query
func (l *Client) PostQueryValidate(query string) ([]Query, error) {
	res := ApiResp{}
	data := []Query{}

	// make the request
	err := l.PostType("text/plain", queryValidateEndpoint, nil, query, &res, &data)

	if err != nil {
		return data, err
	}

	return data, nil
}

// PostQueryValidateSegment checks query interpretation & validates query against existing segments/ schemas
func (l *Client) PostQueryValidateSegment(query string) (interface{}, error) {
	res := ApiResp{}
	djson := json.RawMessage{}
	v := url.Values{}
	v.Set("segments", "true")

	// make the request
	err := l.PostType("text/plain", queryValidateEndpoint, v, query, &res, &djson)
	var data interface{}
	// unmarshal response `data` field into appropriate type depending on response status & message
	// add cases?
	switch res.Status {
	case float64(400):
		if res.Message == "Invalid schema errors" {
			d := []string{}
			e := json.Unmarshal(djson, &d)
			if e != nil {
				return nil, e
			}
			data = d
		}
	}
	if err != nil {
		return data, err
	}
	return data, nil
}
