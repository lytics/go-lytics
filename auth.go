package lytics

import (
	"time"
)

var (
	// Ensure we can write out as a table
	_ TableWriter = (*Auth)(nil)
)

const (
	authEndpoint     = "auth/:id"
	authListEndpoint = "auth"
)

// Auth is the authorization token (apikey, user-token, oauth-token, jwt-key) to provide
// Lytics access to a 3rd party resource.  These are ENCRYPTED and never available to be
// re-read through the api after creation.
type Auth struct {
	Id           string    `json:"id"`
	Aid          int       `json:"aid"`
	Name         string    `json:"name"`
	AccountId    string    `json:"account_id"`
	ProviderId   string    `json:"provider_id"`
	ProviderName string    `json:"provider_name"`
	UserId       string    `json:"user_id"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"updated"`
}

func (m *Auth) Headers() []interface{} {
	return []interface{}{
		"ID", "account_id", "name", "provider", "provider_id", "user_id", "created", "updated",
	}
}
func (m *Auth) Row() []interface{} {
	return []interface{}{
		m.Id, m.AccountId, m.Name, m.ProviderName, m.ProviderId, m.UserId, m.Created.Format(time.RFC3339), m.Updated.Format(time.RFC3339),
	}
}

// GetAuths returns a list of all available auths for an account
// https://learn.lytics.com/documentation/developer/api-docs/auth#auth
func (l *Client) GetAuths() ([]*Auth, error) {
	res := ApiResp{}
	data := []*Auth{}

	// make the request
	err := l.Get(authListEndpoint, nil, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// GetAuth returns a single auth based on id
// https://learn.lytics.com/documentation/developer/api-docs/auth#auth-retrieve-a-single-auth-get
func (l *Client) GetAuth(id string) (Auth, error) {
	res := ApiResp{}
	data := Auth{}

	// make the request
	err := l.Get(parseLyticsURL(authEndpoint, map[string]string{"id": id}), nil, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Other Available Endpoints
// * POST    create auth
// * PUT     update auth
// * DELETE  remove auth
