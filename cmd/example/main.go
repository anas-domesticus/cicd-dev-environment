package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("I'm alive...")
		time.Sleep(2 * time.Second)
	}
}
