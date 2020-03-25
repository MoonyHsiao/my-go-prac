// https://kaviraj.me/understanding-condition-variable-in-go/
package main

import (
	"fmt"
	"sync"
	"time"
)

//Naive way: Use infinite for loop.
type Record struct {
	sync.Mutex
	data string
}

func main() {
	var wg sync.WaitGroup

	rec := &Record{}
	wg.Add(1)
	go func(rec *Record) {
		defer wg.Done()
		for {
			rec.Lock()
			if rec.data != "" {
				fmt.Println("Data:", rec.data)
				rec.Unlock()
				return
			}
			rec.Unlock()
		}
	}(rec)

	time.Sleep(2 * time.Second)
	rec.Lock()
	rec.data = "Naive way: Use infinite for loop."
	rec.Unlock()

	wg.Wait() // wait till all goutine completes
}