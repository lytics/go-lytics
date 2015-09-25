package mock

import (
	"github.com/jarcoal/httpmock"
	"net/http"
)

func RegisterCatalogMocks() {
	// *******************************************************
	// GET SCHEMA FOR TABLE
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/schema/user",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			// if we have a limit param respond with cut down
			if queries.Get("limit") == "2" {
				return httpmock.NewStringResponse(200, readJsonFile("get_catalog_user_limit")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_catalog_user")), nil
		},
	)
}
