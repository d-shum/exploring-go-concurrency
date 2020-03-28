package base

import "sync"

// IncrementAThousandTimes implements naive concurrent realisation without any synchronisation
func IncrementAThousandTimes(valuePointer *int) {
	// WaitGroup guaranties that all funcs will finish but doesnot guarantee that incrementation is safe
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			*valuePointer++
			wg.Done()
		}()
	}
	wg.Wait()
}
