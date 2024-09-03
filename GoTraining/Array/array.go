package main

import (
	"fmt"
)

func main() {
	arr := [4]string{"aa", "bb", "ee", "yy"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
