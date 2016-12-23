package mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jarcoal/httpmock"
	"net/http"
)

func RegisterSegmentMocks() {

	// *******************************************************
	// POST SINGLE SEGMENT
	// *******************************************************
	httpmock.RegisterResponder("POST", "https://api.lytics.io/api/segment",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_segment")), nil
		},
	)

	// *******************************************************
	// GET SINGLE SEGMENT
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/segment/"+MockSegmentID1,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_segment")), nil
		},
	)

	// *******************************************************
	// GET ALL SEGMENTS FOR API KEY
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/segment",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if queries.Get("table") == "content" {
				return httpmock.NewStringResponse(200, readJsonFile("get_content_segments")), nil
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_segments")), nil
		},
	)

	// *******************************************************
	// GET SEGMENT SIZES
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/segment/"+MockSegmentID1+"/sizes",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_segment_size")), nil
		},
	)

	// *******************************************************
	// GET ALL SEGMENT SIZES FOR API KEY
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/segment/sizes",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_segment_sizes")), nil
		},
	)

	// *******************************************************
	// GET SEGMENT ATTRIBUTION FOR 2 SEGMENTS
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/segment/attribution",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			// make sure we have the correct segment ids in the request
			if queries.Get("ids") != MockSegmentID1+","+MockSegmentID2 {
				fmt.Println("Ids Query Failed in Segment Attribution Mock")
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_segment_attribution")), nil
		},
	)

	// *******************************************************
	// GET SEGMENT COLLECTION
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/segmentcollection/"+MockSegmentCollection,
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_segment_collection")), nil
		},
	)

	// *******************************************************
	// GET SEGMENT COLLECTION List
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/segmentcollection",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_segment_collection_list")), nil
		},
	)

	// *******************************************************
	// GET SEGMENT SCAN
	// *******************************************************
	httpmock.RegisterResponder("GET", "https://api.lytics.io/api/segment/f186a334ad7109bbe08880/scan",
		func(req *http.Request) (*http.Response, error) {
			var fail bool

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				fail = true
			}

			// make sure we have the correct segment ids in the request
			if queries.Get("start") == "number2" {
				return httpmock.NewStringResponse(200, readJsonFile("get_segment_scan_2")), nil
			}
			if queries.Get("start") == "number3" {
				return httpmock.NewStringResponse(200, readJsonFile("get_segment_scan_3")), nil
			}

			if fail {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("get_segment_scan_1")), nil
		},
	)

	// *******************************************************
	// POST CREATE SEGMENT
	// *******************************************************
	httpmock.RegisterResponder("POST", "https://api.lytics.io/api/segment",
		func(req *http.Request) (*http.Response, error) {

			body := make(map[string]interface{})
			json.NewDecoder(req.Body).Decode(&body)

			if _, ok := body["segment_ql"]; !ok {
				return httpmock.NewStringResponse(400, readJsonFile("post_segment_create_error")), nil
			}

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("post_segment_create")), nil
		},
	)

	// *******************************************************
	// POST SEGMENT VALIDATE
	// *******************************************************
	httpmock.RegisterResponder("POST", "https://api.lytics.io/api/segment/validate",
		func(req *http.Request) (*http.Response, error) {

			buf := new(bytes.Buffer)
			buf.ReadFrom(req.Body)

			if buf.String() == "" {
				return httpmock.NewStringResponse(400, readJsonFile("post_segment_validate_error")), nil
			}

			queries := req.URL.Query()

			if queries.Get("key") != MockApiKey {
				return httpmock.NewStringResponse(401, readJsonFile("get_error")), nil
			}

			return httpmock.NewStringResponse(200, readJsonFile("post_segment_validate")), nil
		},
	)
}
