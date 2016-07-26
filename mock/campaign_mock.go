package mock

import (
	"github.com/jarcoal/httpmock"
	"net/http"
)

func RegisterCampaignMocks() {
	// *******************************************************
	// GET SINGLE CAMPAIGN
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/program/campaign/"+MockCampaignID,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_campaign")), nil
		},
	)

	// *******************************************************
	// GET CAMPAIGN LIST
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/program/campaign",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if queries.Get("status") == "published" {
				return httpmock.NewStringResponse(200, readJsonFile("get_campaign_list_published")), nil
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_campaign_list")), nil
		},
	)

	// *******************************************************
	// GET SINGLE VARIATION
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/program/campaign/variation/"+MockVariationID,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_variation")), nil
		},
	)

	// *******************************************************
	// GET VARIATION LIST
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/program/campaign/variation",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_variation_list")), nil
		},
	)
}
