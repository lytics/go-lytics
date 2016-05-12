package mock

import (
	"github.com/jarcoal/httpmock"
	"net/http"
)

func RegisterContentMocks() {
	// *******************************************************
	// GET USER CONTENT RECOMMENDATION
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/content/recommend/user/user_id/"+MockUserID,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_user_content_recommendation")), nil
		},
	)
}
