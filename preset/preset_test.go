package preset

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewPreset(t *testing.T) {
	t.Run("success", func(t *testing.T){
		presets, err := NewPreset(ElReto)
		assert.NoError(err)
	})
}
