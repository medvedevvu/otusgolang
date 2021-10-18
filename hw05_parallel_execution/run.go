package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.

func Run(tasks []Task, n, m int) error {
	taskChanal := make(chan Task, len(tasks))
	errChanl := make(chan int, len(tasks))
	okChan := make(chan int, len(tasks))
	errCounter := 0
	okCounter := 0
	for _, item := range tasks {
		taskChanal <- item
	}
	close(taskChanal)
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		go func(inChanal chan Task) {
			defer wg.Done()
			wg.Add(1)
			for el := range inChanal {
				err := el()
				if err != nil {
					errChanl <- 1
				}
				okChan <- 1
			}
		}(taskChanal)
	}
Loop:
	for {
		select {
		case err, hOK := <-errChanl:
			if hOK {
				errCounter += err
			}
		case ok, hOK := <-okChan:
			if hOK {
				okCounter += ok
			}
		default:
			if errCounter > m {
				return ErrErrorsLimitExceeded
			}
			if okCounter == len(tasks) {
				break Loop
			}
			if (errCounter + okCounter) >= len(tasks) {
				break Loop
			}
		}
	}
	wg.Wait()
	return nil
}
