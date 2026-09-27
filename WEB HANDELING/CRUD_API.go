package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main(){
	fmt.Println("Learninh crud ")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	defer res.Body.Close()

	// Check if the request was successful
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Response :", res.Status)
		return
	}
	
	// Read the response body
	data , err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Data :" , string(data))

	// we can convert these data into struct using json.Unmarshal
}