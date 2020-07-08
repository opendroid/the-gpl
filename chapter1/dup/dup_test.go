package dup

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// cd dup
// go test -run TestDuplicateLines -v
func TestDuplicateLines(t *testing.T) {
	assert.Equal(t, true, DuplicateLines())
}
