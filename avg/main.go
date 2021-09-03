package main

import (
	"fmt"
)

func data(number ...int) {
	sum := 0
	length := len(number)
	/*
		for i := 0; i < len(number); i++ {
			sum += (number[i])
		}
		avg := sum / len(number)
		fmt.Println(avg) */
	for _, val := range number {
		// sum += val
		sum = sum + val

	}
	avg := sum / length
	fmt.Println(avg)

}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7}
	data(values...)
}
