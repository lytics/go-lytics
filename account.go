package lytics

import (
	"time"
)

const (
	accountEndpoint     = "account/:id"
	accountListEndpoint = "account"
)

// Account provides the metadata for an account
type Account struct {
	// Unique identifier for this account
	Id string `json:"id"`

	// Fid is an optional user-assigned unique identifier
	Fid string `json:"fid"`

	// Domain is the unique domain associated with this account
	Domain string `json:"domain"`

	// Aid is a simple numeric id for the account
	Aid int `json:"aid"`

	// ParentAid points to the parent Aid if this account is a child
	// otherwise it is the same as Aid
	ParentAid int `json:"parentaid"`

	// ParentId points to the parent Id if this account is a child
	// otherwise it is the same as Id
	ParentId string `json:"parent_id"`

	// Email address for primary contact of this account
	Email string `json:"email"`

	// ApiKey contains the key to use for management API access
	ApiKey string `json:"apikey"`

	// DataApiKey contains the key to use for data read/write access
	DataApiKey string `json:"dataapikey"`

	// Account name
	Name string `json:"name"`

	// Primary time zone for the account
	TimeZone string `json:"timezone,omitempty"`

	// PubUsers defines whether auth is required for entity read requests
	PubUsers bool `json:"pubusers"`

	// WhitelistFields defines the fields which are provided in an entity read
	// request
	WhitelistFields []string `json:"whitelist_fields"`

	// WhitelistDomains defines the list of domains which are able to perform
	// an entity read request
	WhitelistDomains []string `json:"whitelist_domains"`

	// Account creation time
	Created time.Time `json:"created"`

	// Account update time
	Updated time.Time `json:"updated"`
}

// GetAccounts returns a list of all accounts associated with master account
// https://www.getlytics.com/developers/rest-api#accounts-list
func (l *Client) GetAccounts() ([]Account, error) {
	res := ApiResp{}
	data := []Account{}

	// make the request
	err := l.Get(accountListEndpoint, nil, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// GetAccount returns the details of a single account based on id
// https://www.getlytics.com/developers/rest-api#account
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
