package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	//ждем
	<-time.After(duration)
}

func main() {
	fmt.Println("Sleeping for 1 second...")
	sleep(time.Second)
	fmt.Println("Awake!")
}
