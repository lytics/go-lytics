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

			if queries.Get("topics[]") != "mock" || len(queries["topics[]"]) != 2 {
				fail = true
			}

			if queries.Get("visited") != "false" {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_user_content_recommendation")), nil
		},
	)

	// *******************************************************
	// GET SEGMENT CONTENT RECOMMENDATION
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/content/recommend/segment/"+MockSegmentID1,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if queries.Get("rank") != "popular" {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_user_content_recommendation")), nil
		},
	)

	// *******************************************************
	// GET CONTENT DOCUMENTS
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/content/doc",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			if queries.Get("limit") == "1" {
				return httpmock.NewStringResponse(200, readJsonFile("get_document")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_documents")), nil
		},
	)

	// *******************************************************
	// GET TOPIC SUMMARY
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/content/topic/"+MockTopicID,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_topic")), nil
		},
	)
}
