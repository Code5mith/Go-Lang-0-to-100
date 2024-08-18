package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// MARK:  Mutex is also known as mutual exclusion

	// Create a shared counter variable

	count := 0

	// Create a mutex to protect the shared counter
	var mutex sync.Mutex

	// Create multiple goroutines to increment the counter
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				// Lock the mutex to ensure exclusive access to the counter

				// this means when we call the mutex.Lock(), we are locking the shared resource
				// to a single go routine in order to work with that shared resource from a single
				// go routine and after the mutex.Unlock(), will unlock the resource making it accessable to all
				// other go routines.

				mutex.Lock()
				count++
				mutex.Unlock() // Release the lock
			}
		}()
	}

	// Wait for all goroutines to finish
	time.Sleep(time.Second)

	// Print the final count
	fmt.Println("Final count:", count)
}
