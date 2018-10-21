package faker

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaolion/gofaker/random"
)

func TestBoolean(t *testing.T) {
	random.Seed(1000)
	assert.Equal(t, false, Boolean())
	assert.Equal(t, true, Boolean())
	assert.Equal(t, false, Boolean())
}
