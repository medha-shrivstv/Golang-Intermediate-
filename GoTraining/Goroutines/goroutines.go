package main

import "fmt"

func main() {
	go hello(".")
	go hello("x")
	fmt.Println("last line of main")
	// wait
	//Option1 - sleep time
	//Option2 - wait for some input
	x := 0
	fmt.Scan(&x)
	//Option3 - infinite loop
	//for{}
}

func hello(str string) {
	for i := 0; i < 1000; i++ {
		fmt.Print(str)
	}
}
