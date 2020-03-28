package channels

import (
	"sync"
)

// IncrementAThousandTimes implements concurrent incrementation with channels
func IncrementAThousandTimes(valuePointer *int) {
	сhannel := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		// Reaciver guaranties safe incrementation, resceives messages from receiver untill channel is closed
		for val := range сhannel {
			*valuePointer = *valuePointer + val
			if *valuePointer == 1001 {
				close(сhannel)
			}
		}
		wg.Done()
	}()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, сhannel)
	}

	wg.Wait()
	return
}

// Sender sends incrementation value of 1 to receiver
func increment(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	ch <- 1
}
