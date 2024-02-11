# Batch

[![Go Reference](https://pkg.go.dev/badge/github.com/veggiemonk/batch.svg)](https://pkg.go.dev/github.com/veggiemonk/batch)

Split an array/slice into `n` evenly chunks.

Inspired from the blog post by [Paul Di Gian](https://github.com/PaulDiGian) on his blog:
[Split a slice or array in a defined number of chunks in golang](https://pauldigian.com/split-a-slice-or-array-in-a-defined-number-of-chunks-in-golang-but-any-language-really)

**Note**: you might better off just copying the function into your codebase.
It has little code.

See [Go Proverbs](https://go-proverbs.github.io/) for more details.

> A little copying is better than a little dependency.

This library isn't really meant to be imported.
Just copy the one function and adapt it to your needs.
Look at the [tests](batch_test.go) for edge cases.
The benchmarks and fuzzing are just for me to learn and have a playground
to try things out.

<!-- TOC -->

-   [Batch](#batch)
    -   [Installation](#installation)
    -   [Usage](#usage)
    -   [Usage with Cloud Run Jobs](#usage-with-cloud-run-jobs)
    -   [Rationale](#rationale)
    -   [Links](#links)

## Installation

Requires Go 1.18 or later.

Just copy the function in [batch.go](batch.go)

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
    chunks := batch.Slice(s, 3)

    fmt.Println(chunks)
    // length      3       3        4
    // output: [[1 2 3] [4 5 6] [7 8 9 10]]
    // the size of each batch has variation of max 1 item
    // this can spread the load evenly amongst workers
}
```

## Usage with Cloud Run Jobs

```go
batchID = uuid.New().String()
taskCount, _ = strconv.Atoi(os.Getenv("CLOUD_RUN_TASK_COUNT"))
taskIndex, _ = strconv.Atoi(os.Getenv("CLOUD_RUN_TASK_INDEX"))

tt, _ := requestToTasks(request)

batches := batch.Slice(tt, taskCount)
if taskIndex >= len(batches) || taskIndex < 0 {
	return fmt.Errorf("index (%d) out of bounds (max: %d), (id:%s): %w", taskIndex, len(batches), batchID, ErrTaskIndexOutOfBounds)
}

b := batches[taskIndex]
if err := process(b); err != nil {
    return fmt.Errorf("failed to process batch (id:%s): %w", batchID, err)
}
```

## Rationale

Having (almost) same sized batch is useful when you want to distribute the workload evenly across multiple workers.

As opposed to defining the _size of each batch_, we define the _number of batch we want_ to have.

Here a **counter** example:

```go
actions := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
batchSize := 3
batches := make([][]int, 0, (len(actions) + batchSize - 1) / batchSize)

for batchSize < len(actions) {
    actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])
}
batches = append(batches, actions)
fmt.Println(result)
// length       4    |    4    |  2
// output: [[1 2 3 4] [5 6 7 8] [9 10]]
// 2 workers will do double the work of the last worker.
// --> Not what we want.
}
```

This is not ideal when you want to distribute the workload evenly across multiple workers.

The code was taken from [Go wiki - Slice Tricks](https://go.dev/wiki/SliceTricks#batching-with-minimal-allocation).

[//]: # "can be played with here: https://go.dev/play/p/-ULiql4tOTc"

## Links

-   [Split a slice or array in a defined number of chunks in golang](https://pauldigian.com/split-a-slice-or-array-in-a-defined-number-of-chunks-in-golang-but-any-language-really)
-   [Go wiki - Slice Tricks](https://go.dev/wiki/SliceTricks#batching-with-minimal-allocation)
