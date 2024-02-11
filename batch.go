package batch

// Slice evenly slices an array `a` into `n` number of batches
// the size of each batch never deviates more than 1 from the average batch size.
func Slice[T any](a []T, n int) [][]T {
	if n <= 0 {
		return [][]T{}
	}

	batches := make([][]T, 0, n)

	var size, lower, upper int
	l := len(a)

	for i := 0; i < n; i++ {
		lower = i * l / n
		upper = ((i + 1) * l) / n
		size = upper - lower

		a, batches = a[size:], append(batches, a[0:size:size])
	}
	return batches
}

// Deprecated: BatchSlice is an alias function to [Slice] to maintain backward compatibility.
// It waa changed because the name of the package was in the function name,
// which is redundant.
func BatchSlice[T any](a []T, n int) [][]T { return Slice(a, n) }
