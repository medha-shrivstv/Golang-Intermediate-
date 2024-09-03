package main

import "fmt"

func main() {
	ch := make(chan string)
	go send(ch, "a")
	//fmt.Print(ch)
	fmt.Println("waiting in main for channel data ")
	str := <-ch
	fmt.Println("Received ", str)
	fmt.Println("waiting for one for data")
	str = <-ch
	fmt.Println("Received ", str)
	for {
	}

}

func send(ch chan string, str string) {
	//for i := 0; i < 1000; i++ {
	fmt.Println("starting to send data on channel")
	ch <- "send1"
	fmt.Println("after sending data on channel")
	ch <- "send2"
}
