package batch_test

import (
	"fmt"
	"testing"

	"github.com/veggiemonk/batch"
)

func BenchmarkBatchSlice(b *testing.B) {
	for size := 10; size < 1000000; size *= 10 {
		array := genArrayInt(size)
		batchNo := len(array) / 11
		var res [][]int
		b.Run(fmt.Sprint("size=", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res = batch.Slice(array, batchNo)
			}
			b.ReportAllocs()
			_ = res
		})
	}
}

func BenchmarkSliceTrick(b *testing.B) {
	for size := 10; size < 1000000; size *= 10 {
		array := genArrayInt(size)
		batchNo := len(array) / 11
		var res [][]int

		b.Run(fmt.Sprint("size=", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res = sliceTrick(array, batchNo)
			}
			b.ReportAllocs()
			_ = res
		})
	}
}

// sliceTrick is not the same output as Slice but has amazing performance.
// Taken from the Go Wiki in the SliceTricks page
func sliceTrick(actions []int, batchSize int) [][]int {
	if len(actions) == 0 || batchSize < 1 {
		return nil
	}
	batches := make([][]int, 0, (len(actions)+batchSize-1)/batchSize)

	for batchSize < len(actions) {
		actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])
	}
	batches = append(batches, actions)
	return batches
}
