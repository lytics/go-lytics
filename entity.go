package lytics

import (
	"fmt"
	"strings"
)

const (
	entityEndpoint = "entity" // :entitytype/:fieldname/:fieldval, fields
)

// Entity is the main data source for Lytics. All users, content, etc. are referred to as "entities"
type Entity map[string]interface{}

// GetEntity returns all the availble attributes for a single entity (user, content, etc)
// https://www.getlytics.com/developers/rest-api#entity-a-p-i
func (l *Client) GetEntity(entitytype, fieldname, fieldval string, fields []string) (Entity, error) {
	res := ApiResp{}
	data := Entity{}
	toAppend := ""
	endpointParams := map[string]string{}
	params := map[string]string{}

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
		params = map[string]string{
			"fields": strings.Join(fields, ","),
		}
	}

	// make the request
	err := l.Get(parseLyticsURL(endpoint, endpointParams), params, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
