package serial

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func api() bool {
	// 10 ms call
	time.Sleep(1 * time.Millisecond)
	defer wg.Done()
	return true
}

func Init() {
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go api()
	}
	wg.Wait()
	t1 := time.Now()


	fmt.Printf("Concurrent call took %v to run.\n", t1.Sub(t0))

	t3 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		api()
	}
	t4 := time.Now()
	fmt.Printf("Serial call took %v to run.\n", t4.Sub(t3))
}
