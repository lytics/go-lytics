package lytics

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

const (
	entityEndpoint = "entity" // :entitytype/:fieldname/:fieldval, fields
)

// EntityHandler for use in paging
type EntityHandler func(*Entity)

// Entity is the main data source for Lytics. All users, content, etc. are referred to as "entities"
// An Entity is made up of a series of key/value pairs, where the key is a string, and the value may
// be one of a few basic data types, depending on the queries used to create it
type Entity map[string]interface{}

// GetEntity returns all the availble attributes for a single entity (user, content, etc)
// https://www.getlytics.com/developers/rest-api#entity-a-p-i
func (l *Client) GetEntity(entitytype, fieldname, fieldval string, fields []string) (Entity, error) {
	res := ApiResp{}
	data := Entity{}
	toAppend := ""
	endpointParams := map[string]string{}
	params := url.Values{}

	// handle optional endpointParams
	if entitytype != "" {
		toAppend = fmt.Sprintf("/%s", ":entitytype")
		endpointParams["entitytype"] = entitytype

		if fieldname != "" {
			toAppend = fmt.Sprintf("%s/%s", toAppend, ":fieldname")
			endpointParams["fieldname"] = fieldname

			if fieldval != "" {
				toAppend = fmt.Sprintf("%s/%s", toAppend, ":fieldval")
				endpointParams["fieldval"] = fieldval
			}
		}
	}

	// build dynamic endpoint
	endpoint := fmt.Sprintf("%s%s", entityEndpoint, toAppend)

	// if there are also fields, add them
	if len(fields) > 0 {
		params.Add("fields", strings.Join(fields, ","))
	}

	// make the request
	err := l.Get(parseLyticsURL(endpoint, endpointParams), params, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// PrettyJson returns an indented/newlined JSON string representation of the
// entity
func (e *Entity) PrettyJson() string {
	by, _ := json.MarshalIndent(e, "", "  ")
	return string(by)
}
