package mock

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

func RegisterMLMocks() {
	// *******************************************************
	// GET SINGLE ML MODEL
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/ml/"+MockMLID,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_ml_model")), nil
		},
	)

	// *******************************************************
	// GET ALL ML MODELS
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/ml",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_ml_models")), nil
		},
	)
}
