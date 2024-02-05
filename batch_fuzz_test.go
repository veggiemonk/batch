package batch_test

import (
	"testing"

	"github.com/veggiemonk/batch"
)

func FuzzBatchSliceByte(f *testing.F) {
	f.Fuzz(
		func(t *testing.T, a []byte, b int) {
			batch.Slice(a, b)
		},
	)
}
