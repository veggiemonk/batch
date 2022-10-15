# Batch

Split an array/slice into `n` evenly chunks.

Inspired from the blog post by [Paul Di Gian](https://github.com/PaulDiGian) on his blog:
[Split a slice or array in a defined number of chunks in golang](https://pauldigian.com/split-a-slice-or-array-in-a-defined-number-of-chunks-in-golang-but-any-language-really)

## Installation

Requires Go 1.18 or later.

add `github.com/veggiemonk/batch` to your `go.mod` file

then run the following command:

```bash
go mod tidy
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/veggiemonk/batch"
)

func main() {
    s :=  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // Split the slice into 3 even parts
    chunks := batch.BatchSlice(s, 3)

    // Print the chunks
    fmt.Println(chunks)
    // length      3       3        4
    // output: [[1 2 3] [4 5 6] [7 8 9 10]]
    // the size of each batch has variation of max 1 item
    // this can spread the load evenly amongst workers
}
```

## Rationale

Having evenly sized batch is useful when you want to distribute the workload evenly across multiple workers.

As opposed to defining the _size of each batch_, we define the _number of batch we want_ to have.

Here a **counter** example:

```go
package main
import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chunkSize := 3
    var result [][]int
	
	for i := 0; i < len(array); i += chunkSize {
		end := i + chunkSize

		if end > len(array) {
			end = len(array)
		}

		result = append(result, array[i:end])
	}
	
	fmt.Println(result)
	// length       4    |    4    |  2 
	// output: [[1 2 3 4] [5 6 7 8] [9 10]]
	// 2 workers will do double the work of the last worker.
}
```

This is not ideal when you want to distribute the workload evenly across multiple workers.



[//]: # (can be played with here: https://go.dev/play/p/-ULiql4tOTc)

