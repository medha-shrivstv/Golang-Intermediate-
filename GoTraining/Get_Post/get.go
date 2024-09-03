package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	resp, _ := http.Get("https://reqres.in/api/users/2")
	fmt.Println("Respond ", resp)
	fmt.Println("Resp Status = ", resp.Status)
	fmt.Println("Resp Body = ", resp.Body)
	defer resp.Body.Close()
	bodycontent, err := io.ReadAll(resp.Body)
	fmt.Println("BodyContent = ", string(bodycontent))
	fmt.Println(err)
}
