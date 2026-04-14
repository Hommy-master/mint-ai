package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	for i := 0; i < 10000; i++ {
		randomStr := RandString(i)
		assert.Equal(t, i, len(randomStr))
	}

	// -1 panic
}
