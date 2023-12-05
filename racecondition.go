package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter++
		}
	}()

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

//Explanation of Race Condition
//Race Condition Occurrence:

//In this program, two goroutines are incrementing the same counter variable 1000 times each. Since both goroutines access and modify counter without any synchronization mechanism (like mutexes or channels), they might read and write this shared variable simultaneously, leading to unpredictable results.
//Explanation of Why It Occurs:

//The race condition occurs because the operation counter++ is not atomic. It involves three steps: reading the counter value, incrementing it, and writing it back. When these goroutines run concurrently, they might interleave these steps in various ways.
//For example, both goroutines could read the value of counter simultaneously, say 10. Both increment it by 1 separately and write back 11, although counter has been incremented twice. Ideally, it should have been 12.
//The lack of synchronization means that the order of operations across goroutines is non-deterministic, and the final value of counter can vary between different runs of the program.
//Running this program multiple times typically results in different values for the counter, demonstrating the presence of a race condition. You can use the -race flag while running the program (go run -race filename.go) to detect the race condition explicitly.
