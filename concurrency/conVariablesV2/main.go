// https://kaviraj.me/understanding-condition-variable-in-go/
package main

import (
	"fmt"
	"sync"
	"time"
)

// Better way: Use sync.Cond

type Record struct {
	sync.Mutex
	data string

	cond *sync.Cond
}

func NewRecord() *Record {
	r := Record{}
	r.cond = sync.NewCond(&r)
	return &r
}

func main() {
	var wg sync.WaitGroup

	rec := NewRecord()
	wg.Add(1)
	go func(rec *Record) {
		defer wg.Done()
		rec.Lock()
		rec.cond.Wait() //wait會先釋放rec.Lock() 等到Signal之後才繼續
		rec.Unlock()
		fmt.Println("Data:", rec.data)
		return
	}(rec)

	time.Sleep(2 * time.Second)
	rec.Lock()
	rec.data = "Better way: Use sync.Cond"
	rec.Unlock()
	rec.cond.Signal() //塞完值之後再給訊號

	wg.Wait() // wait till all goutine completes
}
