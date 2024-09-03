package main

import (
	"fmt"
)

type Dept struct {
	deptno     int
	dname, loc string
}

func main() {
	depts := [3]Dept{}
	for i := 0; i < 3; i++ {
		fmt.Print("Enter dept no, name and loc  ")
		no1, err := fmt.Scan(&(depts[i].deptno), &(depts[i].dname), &(depts[i].loc))
		fmt.Println(depts[i])
		if no1 != 3 {
			fmt.Println("problem in input, not handling ", err)
		}
	}

	for i := 0; i < len(depts); i++ {
		fmt.Println("Dept: ", depts[i])
	}
}
