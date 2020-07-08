package mas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// cd dup
// go test -run TestCompareNumbers -v
func TestCompareNumbers(t *testing.T) {
	r, d := CompareNumbers(1, 1)
	assert.Equal(t, true, r)
	assert.Equal(t, 0, d)
}
