package mock

import (
	"github.com/jarcoal/httpmock"
	"net/http"
)

func RegisterWorkMocks() {
	// *******************************************************
	// GET SINGLE WORK
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/work/"+MockWorkID,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_work")), nil
		},
	)

	// *******************************************************
	// GET ALL AVAILABLE WORKS FOR API KEY
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/work",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_works")), nil
		},
	)
}
