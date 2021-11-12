package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {
	fmt.Println(time.Now())

	if err := syscall.Nanosleep(&syscall.Timespec{
		Sec:  1,
		Nsec: 0,
	}, nil); err != nil {
		panic(err)
	}

	fmt.Println(time.Now())
}
