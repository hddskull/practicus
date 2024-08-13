package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGaugeMetrics(t *testing.T) {

	t.Run("27 gauge metrics", func(t *testing.T) {
		m := GetGaugeMetrics()
		assert.Equal(t, 27, len(m))
	})
}
