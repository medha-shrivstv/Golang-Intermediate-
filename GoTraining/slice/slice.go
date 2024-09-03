package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	slice1 := primes[0:2]
	fmt.Println("Primes ", primes)
	fmt.Println("Slice1 ", slice1)
	fmt.Println("Length  = ", len(slice1), " and capacity is ", cap(slice1))
	slice1[0] = 111
	slice1[1] = 222
	fmt.Println("Primes ", primes)
	fmt.Println("Slice1 ", slice1)
	fmt.Println("Length  = ", len(slice1), " and capacity is ", cap(slice1))

	slice2 := make([]int, 5)
	fmt.Println(&slice2)
	fmt.Println(slice2)
	slice2[0] = 100
	slice2[3] = 60
	slice2[4] = 11
	fmt.Println(slice2)
	fmt.Println(slice2)
	fmt.Println("Length  = ", len(slice2), " and capacity is ", cap(slice2))
	slice2 = append(slice2, 3)
	fmt.Println(slice2)
	fmt.Println("Length  = ", len(slice2), " and capacity is ", cap(slice2))
	slice2 = append(slice2, 4)
	fmt.Println(slice2)
	slice2 = append(slice2, 4)
	slice2 = append(slice2, 4)
	slice2 = append(slice2, 4)
	fmt.Println("Length  = ", len(slice2), " and capacity is ", cap(slice2))
	slice2 = append(slice2, 5)
	fmt.Println("Length  = ", len(slice2), " and capacity is ", cap(slice2))

}
