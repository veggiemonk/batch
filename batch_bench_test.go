package batch_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/veggiemonk/batch"
)

func BenchmarkBatchSlice(b *testing.B) {
	for size := 10; size < 1000000; size *= 10 {
		array := genArrayInt(size)
		batchNo := len(array) / 11
		res := make([][]int, 0, batchNo)

		b.Run(fmt.Sprint("size=", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res = batch.Slice(array, batchNo)
			}
			b.ReportAllocs()
			_ = res
		})
	}
}

var result [][]int

func BenchmarkSliceCompare(b *testing.B) {
	impls := []struct {
		name string
		fun  func([]int, int) [][]int
	}{
		{
			name: "Trick",
			fun:  sliceTrick[int],
		},
		{
			name: "Batch",
			fun:  batch.Slice[int],
		},
		{
			name: "Simple",
			fun:  simple[int],
		},
	}

	for _, impl := range impls {
		for k := 1; k <= len(prime); k++ {
			size := int(math.Pow(10, float64(k)))
			array := genArrayRandomInt(size)
			batchNo := prime[k-1]

			res := make([][]int, 0, batchNo)

			b.Run(fmt.Sprintf("%s/%d", impl.name, size), func(b *testing.B) {
				b.ReportMetric(float64(batchNo), "#batches")
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					res = impl.fun(array, batchNo)
				}
				result = res
			})
		}
	}
	_ = result
}

var prime = []int{5, 7, 11, 53, 97, 997}

func genArrayRandomInt(n int) []int {
	r := rand.New(rand.NewSource(13378232375))
	a := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		a = append(a, r.Int())
	}
	return a
}

// sliceTrick is not the same output as Slice but has amazing performance.
// Taken from the Go Wiki in the SliceTricks page
func sliceTrick[T any](actions []T, batchSize int) [][]T {
	if len(actions) == 0 || batchSize < 1 {
		return [][]T{}
	}
	batches := make([][]T, 0, (len(actions)+batchSize-1)/batchSize)

	for batchSize < len(actions) {
		actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])
	}
	batches = append(batches, actions)
	return batches
}

// simple is a naive implementation without pre-allocating the resulting slice of slices
func simple[T any](a []T, b int) [][]T {
	var batches [][]T

	l := len(a)

	for i := 0; i < b; i++ {
		low := i * l / b
		upp := ((i + 1) * l) / b

		batches = append(batches, a[low:upp])
	}

	return batches
}
