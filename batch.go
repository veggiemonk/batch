package main

// BatchSlice evenly slices an array `a` into `b` number of batches
// the size of each batch never deviates more than 1 from the average batch size.
func BatchSlice[T any](a []T, b int) [][]T {
	var result [][]T

	l := len(a)

	for i := 0; i < b; i++ {
		min := i * l / b
		max := ((i + 1) * l) / b

		result = append(result, a[min:max])
	}

	return result
}
