package main

import (
	"fmt"

	"github.com/d-shum/exploring-go-concurrency/base"
	"github.com/d-shum/exploring-go-concurrency/channels"
)

func main() {
	value := 1
	channels.IncrementAThousandTimes(&value)
	fmt.Println(value)
	value = 1
	base.IncrementAThousandTimes(&value)
	fmt.Println(value)
}
