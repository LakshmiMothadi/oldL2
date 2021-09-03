package main

import "fmt"

func compareTriplets(a [3]int, b [3]int) [2]int {

	A, B := 0, 0
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			A++
		} else if b[i] > a[i] {
			B++
		}

	}

	return [2]int{A, B}
}

func main() {
	var a, b [3]int
	fmt.Println("a:")
	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Println("b:")
	fmt.Scan(&b[0], &b[1], &b[2])

	result := compareTriplets(a, b)
	fmt.Println(result)
}


/*fmt.Println("a:")
	fmt.Scanf("%d", &a[0])
	fmt.Scanf("%d", &a[1])
	fmt.Scanf("%d", &a[2])

	fmt.Println("b:")
	fmt.Scanf("%d", &b[0])
	fmt.Scanf("%d", &b[1])
	fmt.Scanf("%d", &b[2])
	*/