package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res,err:=http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err !=nil{
		fmt.Println("Error")
		return
	}

	defer res.Body.Close()
	data , err :=ioutil.ReadAll(res.Body)
	if err!=nil{
		return
	}

	fmt.Println(string(data))

}