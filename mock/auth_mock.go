package mock

import (
	"github.com/jarcoal/httpmock"
	"net/http"
)

func RegisterAuthMocks() {
	// *******************************************************
	// GET SINGLE AUTH
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/auth/"+MockAuthID,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_auth")), nil
		},
	)

	// *******************************************************
	// GET ALL AUTHS FOR API KEY
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/auth",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_auths")), nil
		},
	)
}
