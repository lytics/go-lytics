package lytics

import (
	"net/url"
)

const (
	schemaEndpoint               = "schema/:table"           // limit
	schemaTableFieldinfoEndpoint = "schema/:table/fieldinfo" // limit, fields
	routeSchemaStreams           = "schema/_streams"
)

/*

{
  "data": [
    {
      "hidden": false,
      "ct": 24,
      "curct": 0,
      "stream": "cm_content",
      "last_msg_ts": "1482433744000",
      "last_update_ts": "1483486363322",
      "metrics": [
        {
          "ct": 0,
          "ts": "1483471963000"
        },
        {
          "ct": 0,
          "ts": "1483477363000"
        },
        {
          "ct": 0,
          "ts": "1483482763000"
        }
      ],
      "recent_events": [
        {
          "ContentUrl": [
            "http://createsend.com/t/t-2CD8ACA1878CC54C/t"
          ],
          "Html": [

*/

type (
	Schema struct {
		Name     string   `json:"name"`
		ByFields []string `json:"by_fields"`
		Columns  Columns  `json:"columns"`
	}

	Columns []Column
	Column  struct {
		As         string   `json:"as"`
		IsBy       bool     `json:"is_by"`
		Type       string   `json:"type"`
		ShortDesc  string   `json:"shortdesc"`
		LongDesc   string   `json:"longdesc"`
		Froms      []string `json:"froms"`
		Identities []string `json:"identities"`
	}
	Stream struct {
		Name           string       `json:"stream"`
		Ct             int          `json:"ct"`
		LastMsgTime    JsonTime     `json:"last_msg_ts,omitempty"`
		LastUpdateTime JsonTime     `json:"last_update_ts,omitempty"`
		Recent         []url.Values `json:"recent_events,omitempty"`
	}
)

// GetSchema returns the data schema for an account
// https://www.getlytics.com/developers/rest-api#schema
func (l *Client) GetSchema(table string) (Schema, error) {
	res := ApiResp{}
	data := Schema{}

	// make the request
	err := l.Get(parseLyticsURL(schemaEndpoint, map[string]string{"table": table}), nil, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// GetStreams returns the data stream field introspection for an account
// https://www.getlytics.com/developers/rest-api#streams
func (l *Client) GetStreams(stream string) ([]*Stream, error) {
	res := ApiResp{}
	data := []*Stream{}

	// make the request
	err := l.Get(parseLyticsURL(routeSchemaStreams, nil), nil, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
