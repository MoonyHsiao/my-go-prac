// https://kaviraj.me/understanding-condition-variable-in-go/
package main

import (
	"fmt"
	"sync"
	"time"
)

// Better way: Use chan

type Record struct {
	sync.Mutex
	data string
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan struct{})

	rec := &Record{}
	wg.Add(1)
	go func(rec *Record) {
		defer wg.Done()
		<-ch
		fmt.Println("Data:", rec.data)
		return
	}(rec)

	time.Sleep(2 * time.Second)
	rec.Lock()
	rec.data = "Better way: Use chan"
	rec.Unlock()
	ch <- struct{}{}

	wg.Wait() // wait till all goutine completes
}