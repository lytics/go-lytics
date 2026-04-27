package lytics

import (
	"github.com/stretchr/testify/assert"
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
