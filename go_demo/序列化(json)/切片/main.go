package main

import (
	"encoding/json"
	"fmt"
)

func testMap() {
	//Map := make(map[string]interface{})
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "hah"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(string(data))
}

func main() {
	testMap()
}
