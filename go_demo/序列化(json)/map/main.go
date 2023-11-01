package main

import (
	"encoding/json"
	"fmt"
)

func testSlice() {
	var slice []map[string]interface{}
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "hah"
	m["age"] = 32

	slice = append(slice, m)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}

func main() {
	testSlice()
}
