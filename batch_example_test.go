package batch_test

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/veggiemonk/batch"
)

func Example() {
	batchID := "38529598-becf-4bc6-8cfd-7308ba8e22f3"
	// to mimic Cloud Run Jobs behavior
	os.Setenv("CLOUD_RUN_TASK_COUNT", "3")
	os.Setenv("CLOUD_RUN_TASK_INDEX", "1")
	if err := run(batchID); err != nil {
		panic(err)
	}
	// Output:
	// [two three]
}

var (
	request                 = "request"
	ErrNoTaskFound          = errors.New("task not found")
	ErrTaskIndexOutOfBounds = errors.New("task index out of bounds")
)

func run(batchID string) error {
	tasks, err := requestToTasks(request)
	if err != nil {
		return fmt.Errorf("request list tasks (id:%s): %w", batchID, err)
	}
	b, err := CloudRunJobs(tasks)
	if err := process(b); err != nil {
		return fmt.Errorf("process batch (id:%s): %w", batchID, err)
	}
	return nil
}

func requestToTasks(req string) ([]string, error) {
	return []string{"one", "two", "three", "four", "five"}, nil
}

func process(tasks []string) error {
	fmt.Printf("%v\n", tasks)
	return nil
}

func CloudRunJobs[T any](tasks []T) ([]T, error) {
	if len(tasks) == 0 {
		return nil, ErrNoTaskFound
	}

	taskCount, err := strconv.Atoi(os.Getenv("CLOUD_RUN_TASK_COUNT"))
	if err != nil {
		return nil, err
	}
	taskIndex, err := strconv.Atoi(os.Getenv("CLOUD_RUN_TASK_INDEX"))
	if err != nil {
		return nil, err
	}

	batches := batch.Slice(tasks, taskCount)
	if taskIndex >= len(batches) || taskIndex < 0 {
		return nil, fmt.Errorf("index (%d) out of bounds (max: %d): %w", taskIndex, len(batches), ErrTaskIndexOutOfBounds)
	}

	b := batches[taskIndex]
	return b, nil
}
