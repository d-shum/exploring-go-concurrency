package atomic

import (
	"sync"
	"sync/atomic"
)

// IncrementAThousandTimes implemented with atomic incrementation
func IncrementAThousandTimes(valuePointer *int64) {

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// atomic write to addres
			atomic.AddInt64(valuePointer, 1)
			wg.Done()
		}()
	}
	wg.Wait()
}
