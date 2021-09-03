package main

import (
	"fmt"
)

func data() {
	var slice []string
	slice = append(slice, "added")
	//slice[0] = "extra"
	fmt.Println(slice)
}
func anotherData(){
	mapping := make(map[string]int)
	mapping["one"] = 1
	mapping["two"] = 2
	fmt.Println(mapping["two"])
	secondMapping := map[string]string{
		"myName": "lakshmi",
		"area": "kadapa",

	}
	fmt.Println(secondMapping)


}
func main() {
	data()
	anotherData()
}
