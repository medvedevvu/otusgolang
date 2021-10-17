package hw05parallelexecution

import (
	"errors"
	"fmt"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	curOkCounter , curErrorCounter  := 0 , 0
	var (
		errChan chan struct{}
		okChan  chan struct{}
		alarmStop chan struct{}
	)

   defer func() {
	   close(okChan)
	   close(errChan)
	   close(alarmStop)
   }()

	for  {
		if len(tasks) == 0 {
		  break
		}
		//1. Возьмем n заданий
		curTasks := make([]Task, n, n)
		copy(curTasks, tasks[:n])
		tasks = tasks[:n]
		errChan   = make(chan struct{}, n)
		okChan    = make(chan struct{}, n)
		alarmStop = make(chan struct{}, n)
		fmt.Println( len(tasks) , len(curTasks) )
		for _, task := range curTasks {
			go func(t Task) {
				err := task()
				if err != nil {
					errChan <- struct{}{}
					return
				}
				okChan <- struct{}{}
				return
			}(task)
		}
	}

	for  {
 	  if len(tasks) == 0 {
			break
	     }
 	  select {
	    case _, ok := <-okChan:
			 if ok {
				 curOkCounter++
			 }
	    case _, ok := <-errChan:
			if ok {
				curErrorCounter++
			}
	    default:
		    if curErrorCounter == m {
			// Побить все задания
			var i int
			for {
				if i == n {
					break
				}
				alarmStop <- struct{}{}
				i++
			}
			return ErrErrorsLimitExceeded
		    }
	    }
	}
	return nil
}
