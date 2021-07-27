package lytics

import (
	"time"
)

var (
	// Ensure we can write out as a table
	_ TableWriter = (*AccountUser)(nil)
)

const (
	userEndpoint     = "user/:id"
	userListEndpoint = "user"
)

// AccountUser describes a user that has access to an account.
type AccountUser struct {
	Id         string          `json:"id"`
	Email      string          `json:"email"`
	Name       string          `json:"name"`
	AuthType   string          `json:"auth_type"`
	Password   string          `json:"password"`
	AccountId  string          `json:"account_id"`
	Roles      []string        `json:"roles"`
	Accounts   []AccountAccess `json:"accounts"`
	Created    time.Time       `json:"created"`
	Updated    time.Time       `json:"updated"`
	LastLogon  time.Time       `json:"last_logon"`
	IsSysadmin bool            `json:"is_sysadmin"`
}

// AccountAccess describes relationship between user and that account (roles)
type AccountAccess struct {
	AccountId string    `json:"account_id"`
	Roles     []string  `json:"roles"`
	UserId    string    `json:"granted_by"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
}

func (m *AccountUser) setup(acct *Account) {
	m.AccountId = acct.Id
	for _, aa := range m.Accounts {
		if aa.AccountId == m.AccountId {
			m.Roles = aa.Roles
		}
	}
}

func (m *AccountUser) Headers() []interface{} {
	return []interface{}{
		"ID", "account_id", "email", "name", "roles", "last_logon", "created", "updated",
	}
}
func (m *AccountUser) Row() []interface{} {
	return []interface{}{
		m.Id, m.AccountId, m.Email, m.Name, m.Roles, m.LastLogon, m.Created.Format(time.RFC3339), m.Updated.Format(time.RFC3339),
	}
}

// GetUser returns a single user
// https://learn.lytics.com/documentation/developer/api-docs/user#auth
func (l *Client) GetUser(id string) (AccountUser, error) {
	res := ApiResp{}
	data := AccountUser{}

	acct, err := l.GetAccount("current")
	if err != nil {
		// hm......
		return data, err
	}

	// make the request
	err = l.Get(parseLyticsURL(userEndpoint, map[string]string{"id": id}), nil, nil, &res, &data)
	if err != nil {
		return data, err
	}
	data.setup(&acct)
	return data, nil
}

// GetUsers returns a list of all users
// https://learn.lytics.com/documentation/developer/api-docs/user#user-list-user-list-get
func (l *Client) GetUsers() ([]*AccountUser, error) {
	res := ApiResp{}
	data := []*AccountUser{}

	acct, err := l.GetAccount("current")
	if err != nil {
		// hm......
		return data, err
	}

	// make the request
	err = l.Get(userListEndpoint, nil, nil, &res, &data)
	if err != nil {
		return data, err
	}

	for _, user := range data {
		user.setup(&acct)
	}

	return data, nil
}

// Other Available Endpoints
// * POST    create user
// * PUT     update user
// * DELETE  remove user
// * POST    set new password for user
