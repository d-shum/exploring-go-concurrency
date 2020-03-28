package lock

import "sync"

// IncrementAThousandTimes implemented with mutex for access synchronisation
func IncrementAThousandTimes(valuePointer *int) {

	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			*valuePointer++
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}
