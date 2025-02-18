package serial

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func api(i int) bool {
	// 10 ms call
	time.Sleep(10 * time.Millisecond)
	if i%100 == 0 {
		fmt.Println(i)
	}
	defer wg.Done()
	return true
}

func Init() {
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go api(i)
	}
	t1 := time.Now()
	wg.Wait()

	fmt.Printf("Concurrent call took %v to run.\n", t1.Sub(t0))

	t3 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		api(i)
	}
	t4 := time.Now()
	fmt.Printf("Serial call took %v to run.\n", t4.Sub(t3))
}
