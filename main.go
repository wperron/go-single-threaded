package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"syscall"
	"time"
)

func init() {
	runtime.GOMAXPROCS(1)
	debug.SetMaxThreads(5)
}

func main() {
	fmt.Println(time.Now())

	wg := new(sync.WaitGroup)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			if err := syscall.Nanosleep(&syscall.Timespec{
				Sec:  1,
				Nsec: 0,
			}, nil); err != nil {
				panic(err)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(time.Now())
}
