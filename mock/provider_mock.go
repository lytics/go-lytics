package mock

import (
	"github.com/jarcoal/httpmock"
	"net/http"
)

func RegisterProviderMocks() {
	// *******************************************************
	// GET SINGLE PROVIDER
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/provider/"+MockProviderID,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_provider")), nil
		},
	)

	// *******************************************************
	// GET ALL AVAILABLE PROVIDERS
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/provider",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_providers")), nil
		},
	)
}
