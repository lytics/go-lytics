package lytics

import (
	"time"
)

var (
	// Ensure we can write out as a table
	_ TableWriter = (*Account)(nil)
)

const (
	accountEndpoint     = "account/:id"
	accountListEndpoint = "account"
)

// Account represents a Lytics Account.  For customers that have more than one
// Account it is probably an "Environment" (production, test, dev, etc).
type Account struct {
	Id               string    `json:"id"`
	Created          time.Time `json:"created"`
	Updated          time.Time `json:"updated"`
	Fid              string    `json:"fid"`
	Domain           string    `json:"domain"`
	ParentId         string    `json:"parent_id"`
	PartnerId        string    `json:"partner_id"`
	PackageId        string    `json:"package_id"`
	Email            string    `json:"email"`
	ApiKey           string    `json:"apikey"`
	DataApiKey       string    `json:"dataapikey"`
	Name             string    `json:"name"`
	TimeZone         string    `json:"timezone,omitempty"`
	PubUsers         bool      `json:"pubusers"`
	WhitelistFields  []string  `json:"whitelist_fields"`
	WhitelistDomains []string  `json:"whitelist_domains"`
}

func (m *Account) Headers() []interface{} {
	return []interface{}{
		"ID", "domain", "name", "email", "created", "updated",
	}
}
func (m *Account) Row() []interface{} {
	return []interface{}{
		m.Id, m.Domain, m.Name, m.Email, m.Created.Format(time.RFC3339), m.Updated.Format(time.RFC3339),
	}
}

// GetAccounts returns a list of all accounts associated with master account
// https://learn.lytics.com/api-docs/account#accounts-list
func (l *Client) GetAccounts() ([]*Account, error) {
	res := ApiResp{}
	var data []*Account

	// make the request
	err := l.Get(accountListEndpoint, nil, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// GetAccount returns the details of a single account based on id
// http://learn.lytics.com/api-docs/account#account
func (l *Client) GetAccount(id string) (Account, error) {
	res := ApiResp{}
	data := Account{}

	// make the request
	err := l.Get(parseLyticsURL(accountEndpoint, map[string]string{"id": id}), nil, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Other Available Endpoints
// * POST    create account
// * PUT     upate account
// * DELETE  remove account
