package lytics

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/araddon/gou"
)

const (
	entityEndpoint = "entity"
)

// EntityHandler for use in paging
type EntityHandler func(*Entity)

// Entity is the main data source for Lytics User Profiles/Entities.
// All users, content, etc. are referred to as "entities".
type Entity struct {
	Fields map[string]interface{}
	Meta   map[string]interface{}
}

// GetEntity returns all the availble attributes for a single entity (user, content, etc)
// https://learn.lytics.com/api-docs/personalization
func (l *Client) GetEntity(entitytype, fieldname, fieldval string, fields []string) (*Entity, error) {
	return l.GetEntityParams(entitytype, fieldname, fieldval, fields, url.Values{})
}

// GetEntity returns all the availble attributes for a single entity (user, content, etc)
// https://learn.lytics.com/api-docs/personalization
func (l *Client) GetEntityParams(entitytype, fieldname, fieldval string, fields []string, params url.Values) (*Entity, error) {
	res := ApiResp{}
	data := Entity{}
	toAppend := ""
	endpointParams := map[string]string{}

	// handle optional endpointParams
	if entitytype == "" {
		entitytype = "user"
	}

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
	// build dynamic endpoint
	endpoint := fmt.Sprintf("%s%s", entityEndpoint, toAppend)

	// if there are also fields, add them
	if len(fields) > 0 {
		params.Add("fields", strings.Join(fields, ","))
	}

	// make the request
	err := l.Get(parseLyticsURL(endpoint, endpointParams), params, nil, &res, &data.Fields)
	if err != nil {
		gou.Errorf("wtf %v", err)
		return nil, err
	}
	data.Meta = res.Meta
	return &data, nil
}

func (e *Entity) PrettyJson() string {
	by, _ := json.MarshalIndent(e.Fields, "", "  ")
	return string(by)
}
