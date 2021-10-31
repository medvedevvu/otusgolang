package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	var (
		errCounter int32
		okCounter  int32
		taskChanal chan Task
	)
	taskChanal = make(chan Task)
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range taskChanal {
				if err := task(); err != nil {
					atomic.AddInt32(&errCounter, 1)
				} else {
					atomic.AddInt32(&okCounter, 1)
				}
			}
		}()
	}
	for _, task := range tasks {
		if atomic.LoadInt32(&errCounter) >= int32(m) {
			break
		}
		taskChanal <- task
	}
	close(taskChanal)
	wg.Wait()
	if errCounter >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
