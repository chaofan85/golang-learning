package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range arr {
		if num%2 == 0 {
			// fmt.Printf("%v is even.\n", num)
			fmt.Println(num, "is even.")
		} else {
			// fmt.Printf("%v is odd.\n", num)
			fmt.Println(num, "is odd.")
		}
	}
}
