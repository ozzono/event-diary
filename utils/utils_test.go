package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustom(t *testing.T) {
	custom, err := Custom("12:12")
	assert.NoError(t, err)
	t.Logf("custom time %s", custom)
}
