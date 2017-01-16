package lytics

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/bmizerany/assert"
)

func TestJsonFlatten(t *testing.T) {
	msg := `{
		"reach":4,
		"source":"http://twitter.com/tweetbutton",
		"tweet_id":"302601869862252544",
		"bigint":7000000000000000000000,
		"twuser_id":"829488326",
		"twuser_location":"",
		"twuser_screenname":"maranda_hampton",
		"url":["https://www.athletepath.com/join"],
		"nested": {
			"a": [
				{
					"b": {
						"c": [
							[
								{
									"d": "1"
								}
							]
						]
					}
				}
			]
		},
		"heterogeneousList": ["a", 3.14],
		"hasNullVal": null
	}`

	flat, err := FlattenJson([]byte(msg))
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(flat["reach"]))
	assert.Equal(t, "4", flat["reach"][0])
	assert.Equal(t, "http://twitter.com/tweetbutton", flat["source"][0])
	assert.Equal(t, "302601869862252544", flat["tweet_id"][0])
	assert.Equal(t, "829488326", flat["twuser_id"][0])
	assert.Equal(t, "7000000000000000000000", flat["bigint"][0])
	assert.Equal(t, "", flat["twuser_location"][0])
	assert.Equal(t, "maranda_hampton", flat["twuser_screenname"][0])
	assert.Equal(t, "https://www.athletepath.com/join", flat["url"][0])
	assert.Equal(t, "1", flat["nested.a[0].b.c[0][0].d"][0])
	assert.Equal(t, "a", flat["heterogeneousList"][0])
	assert.Equal(t, "3.14", flat["heterogeneousList"][1])
	_, ok := flat["hasNullVal"]
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(flat["hasNullVal"]))

	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(msg), &jsonMap)
	assert.Tf(t, err == nil, "must not err: %v", err)
	qs := make(url.Values)
	err = FlattenJsonMapIntoQs(qs, jsonMap, "")
	assert.Tf(t, err == nil, "must not err: %v", err)
	assert.Tf(t, qs.Get("bigint") == "7000000000000000000000", "must not have exponent format")
}
