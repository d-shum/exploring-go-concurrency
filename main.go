package main

import (
	"fmt"

	"github.com/d-shum/exploring-go-concurrency/base"
	"github.com/d-shum/exploring-go-concurrency/atomic"
	"github.com/d-shum/exploring-go-concurrency/channels"
)

func main() {
	value := 1
	channels.IncrementAThousandTimes(&value)
	fmt.Println(value)
	value = 1
	base.IncrementAThousandTimes(&value)
	fmt.Println(value)
	value64 := int64(1)
	atomic.IncrementAThousandTimes(&value64)
	fmt.Println(value64)
}
