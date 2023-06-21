package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fairblock/dkg-core/x/tss/types"
)

func TestDefaultParams(t *testing.T) {
	params := types.DefaultParams()

	assert.NoError(t, params.Validate())
}
