package main

import (
	"fmt"
	"strconv"
)

func main() {

	ch := make(chan string, 3)
	go sender(ch, "a")
	go receiver(ch)
	for {
	}

}
func receiver(ch chan string) {
	for str := range ch {
		fmt.Println("Received ", str)
	}
	fmt.Println("Out of for loop in Receiver")
}
func sender(ch chan string, str string) {
	for i := 0; i < 15; i++ {
		fmt.Println("starting to send data on channel ", i)
		n1 := strconv.Itoa(i)
		ch <- "send" + n1
	}
	close(ch)
}
