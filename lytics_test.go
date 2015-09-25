package lytics

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestParseLyticsURL(t *testing.T) {
	str := "this/:id/a/:value"
	params := map[string]string{
		"id":    "is",
		"value": "test",
	}

	res := parseLyticsURL(str, params)

	assert.Equal(t, res, "this/is/a/test")
}
