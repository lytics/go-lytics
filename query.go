package lytics

import (
	"net/url"
	"time"
)

const (
	queryEndpoint     = "query/:id"
	queryListEndpoint = "query"
	queryTestEndpoint = "query/_test"
)

type Query struct {
	Id      string                `json:"id"`
	Created time.Time             `json:"created"`
	Updated time.Time             `json:"updated"`
	Alias   string                `json:"alias"`
	Table   string                `json:"table"`
	Text    string                `json:"text"`
	Fields  map[string]QueryField `json:"fields"`
}
type QueryField struct {
	As         string   `json:"as"`
	IsBy       bool     `json:"is_by"`
	Type       string   `json:"type"`
	ShortDesc  string   `json:"shortdesc"`
	LongDesc   string   `json:"longdesc"`
	Identities []string `json:"identities"`
	Froms      []string `json:"froms"`
}

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
// https://www.getlytics.com/developers/rest-api#segment-sizes
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

// Other Available Endpoints
// * POST    create account
// * PUT     upate account
// * DELETE  remove account
