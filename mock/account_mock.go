package mock

import (
	"github.com/jarcoal/httpmock"
	"net/http"
)

func RegisterAccountMocks() {
	// *******************************************************
	// GET SINGLE ACCOUNT
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/account/"+MockAccountID,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_account")), nil
		},
	)

	// *******************************************************
	// GET ALL ACCOUNTS FOR API KEY
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/account",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_accounts")), nil
		},
	)
}
