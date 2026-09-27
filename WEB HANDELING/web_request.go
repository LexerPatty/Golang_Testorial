package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	fmt.Println("Web Request Handling")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1 ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response Status: %T\n",resp )
	fmt.Println("Response", resp)

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Response Body: ", string(data))
}