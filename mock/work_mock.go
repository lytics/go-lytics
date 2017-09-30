package mock

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jarcoal/httpmock"
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

	// *******************************************************
	// CREATE A NEW WORK
	// *******************************************************
	httpmock.RegisterResponder("POST", "https://api.lytics.io/api/work",
		func(req *http.Request) (*http.Response, error) {
			// if invalid key return unauthorized
			queries := req.URL.Query()
			if queries.Get("key") != MockApiKey {
				return httpmock.NewStringResponse(401, readJsonFile("post_work_fail")), nil
			}

			// if trying to create duplicate work
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				panic("failed to get body of mock")
			}
			if strings.Contains(string(body), "iamaduplicatesegmentandshouldfail") {
				return httpmock.NewStringResponse(400, readJsonFile("post_duplicate_work")), nil
			}

			// if success
			return httpmock.NewStringResponse(201, readJsonFile("post_work")), nil
		},
	)
}
