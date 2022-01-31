// to represent race condition, when stingy adds 10 and spendy substracts 10
// at the same time

package main

import (
	"sync/atomic"
	"time"
)

var (
	money int32 = 100
)

func stingy() {
	for i := 1; i <= 1000; i++ {

		atomic.AddInt32(&money, 10)

		time.Sleep(1 * time.Millisecond)
	}
	println("stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(&money, -10)
		time.Sleep(1 * time.Millisecond)
	}
	println("spendy Done")
}

func main() {
	// both are initialised as threads, as 'go' is mentioned
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	print(money)
}

// this program should give wrong output (not 100), but idk why it's not
// race condition should occur, but it's working fine.

// atomic variables can be used to replce mutex locks
