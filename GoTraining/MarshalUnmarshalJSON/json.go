package main

import (
	"encoding/json"
	"fmt"
)

type Dept struct {
	Deptno int
	Dname  string
	Loc    string
}

func main() {
	//func Marshal(v any) ([]byte, error)
	dept := Dept{10, "HR", "Hyd"}
	barr, err := json.Marshal(dept)
	fmt.Println(string(barr), err)
	//func Unmarshal(data []byte, v any) error
	str := `{"Deptno":10,"Dname":"HR","Loc":"Hyd"}`
	d1 := Dept{}
	err = json.Unmarshal([]byte(str), &d1)
	fmt.Println(d1, err)
}
