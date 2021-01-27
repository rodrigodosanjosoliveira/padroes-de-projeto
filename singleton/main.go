package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		go getInstance()
	}

	fmt.Scanln()
}