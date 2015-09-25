package mock

import (
	"github.com/jarcoal/httpmock"
	"net/http"
)

func RegisterUserMocks() {
	// *******************************************************
	// GET SINGLE USER
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/user/"+MockUserID,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_user")), nil
		},
	)

	// *******************************************************
	// GET ALL USERS FOR API KEY
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/user",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_users")), nil
		},
	)
}
