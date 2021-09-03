package main
import "fmt"

func compareTriplets(a []int, b []int) []int {
	fmt.Println("a:")
	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Println("b:")
	fmt.Scan(&b[0], &b[1], &b[2])
	A, B := 0, 0
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			A++
		} else if b[i] > a[i] {
			B++
		}

	}

	return []int{A, B}
}

func main() {
	var a, b []int
	result := compareTriplets(a, b)
	fmt.Println(result)
}


