package main

import (
	"fmt"
)

func main() {
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}
	close(charChannel)

	for s := range charChannel {
		fmt.Print(s)
	}
}
