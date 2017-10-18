package mock

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

func RegisterFaceookMocks() {
	// *******************************************************
	// GET LOOKALIKE STATUS
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/provider/facebook/lookalike/status",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, readJsonFile("get_lookalike_status")), nil
		},
	)

	// *******************************************************
	// CREATE LOOKALIKE
	// *******************************************************
	httpmock.RegisterResponder("POST", "https://api.lytics.io/api/provider/facebook/lookalike/create",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, readJsonFile("create_lookalike")), nil
		},
	)
}
