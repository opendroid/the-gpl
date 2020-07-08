package mas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// cd dup
// go test -run TestLearnStringerInterface -v
func TestLearnStringerInterface(t *testing.T) {
	learnStringerInterface()
	assert.Equal(t, true, true)
}
