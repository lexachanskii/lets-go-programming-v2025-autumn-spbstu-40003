package table

import (
	"container/heap"
	"errors"
	"fmt"

	myHeap "github.com/lexachanskii/task-2-2/internal/heap"
)

var (
	errOutOfRange = errors.New("favourite dish number is bigger than dishes list length")
)

func inputHeap(h *myHeap.HeapInt) (int, error) {
	var (
		dishes    int
		favourite int
		dishNum   int
	)

	if _, err := fmt.Scan(&dishes); err != nil {
		return 0, fmt.Errorf("error while reading dishes count: %w", err)
	}

	for range dishes {
		if _, err := fmt.Scan(&dishNum); err != nil {
			return 0, fmt.Errorf("error while reading dishes list: %w", err)
		}

		heap.Push(h, dishNum)
	}

	if _, err := fmt.Scan(&favourite); err != nil {
		return 0, fmt.Errorf("error while reading favourite dish: %w", err)
	}

	return favourite, nil
}

func getFavourite(favourite int, h *myHeap.HeapInt) (int, error) {
	if h.Len() < favourite {
		return -1, fmt.Errorf("lenth error %w", errOutOfRange)
	}
	for range favourite - 1 {
		heap.Pop(h)
	}

	return heap.Pop(h).(int), nil
}

func Table() (int, error) {
	var heap myHeap.HeapInt

	favourite, err := inputHeap(&heap)
	if err != nil {
		return -1, fmt.Errorf("error in heap input %w", err)
	}

	bestDish, err := getFavourite(favourite, &heap)
	if err != nil {
		return -1, fmt.Errorf("error while getting favourite dish %w", err)
	}

	return bestDish, nil
}
