package mock

import (
	"github.com/jarcoal/httpmock"
	"net/http"
)

func RegisterEntityMocks() {
	// *******************************************************
	// GET ENTITY
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/entity/user/email/sample@sample.com",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if queries.Get("fields") == "score_intensity" {
				return httpmock.NewStringResponse(200, readJsonFile("get_entity_fields")), nil
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_entity")), nil
		},
	)
}
