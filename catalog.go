package lytics

const (
	schemaEndpoint               = "schema/:table"           // limit
	schemaTableFieldinfoEndpoint = "schema/:table/fieldinfo" // limit, fields
)

type Schema struct {
	Name     string   `json:"name"`
	ByFields []string `json:"by_fields"`
	Columns  Columns  `json:"columns"`
}

type Columns []Column
type Column struct {
	As         string   `json:"as"`
	IsBy       bool     `json:"is_by"`
	Type       string   `json:"type"`
	ShortDesc  string   `json:"shortdesc"`
	LongDesc   string   `json:"longdesc"`
	Froms      []string `json:"froms"`
	Identities []string `json:"identities"`
}

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
