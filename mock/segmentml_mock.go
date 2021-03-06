package mock

import (
	"github.com/jarcoal/httpmock"
	"net/http"
)

func RegisterSegmentMLMocks() {
	// *******************************************************
	// GET SINGLE SEGMENTML MODEL
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/segmentml/"+MockSegmentMLID,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_segmentml")), nil
		},
	)

	// *******************************************************
	// GET ALL SEGMENTML MODELS
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/segmentml",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_segmentml_models")), nil
		},
	)
}
