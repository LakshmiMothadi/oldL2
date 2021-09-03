package main

import (
	"fmt"
)

func compareTriplets(a [3]int32, b [3]int32) [2]int32 {
	fmt.Println("a:")
	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Println("b:")
	fmt.Scan(&b[0], &b[1], &b[2])
	var A, B int32 = 0, 0
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			A++
		} else if b[i] > a[i] {
			B++
		}

	}

	return [2]int32{A, B}
}

func main() {
	var a, b [3]int32
	result:= compareTriplets(a, b)
	fmt.Println(result)
}
