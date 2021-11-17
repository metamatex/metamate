package testing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStdT(t *testing.T) {
	stdT := StdT{t: t}

	stdT.Run("case a", func(t TB) {
		t.Parallel()

		assert.True(t, true)
	})

	stdT.Run("case b", func(t TB) {
		t.Parallel()

		assert.True(t, true)
	})
}