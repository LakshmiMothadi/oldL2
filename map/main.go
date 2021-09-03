package main
import (
	"fmt"
)
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
	anotherData()
}
