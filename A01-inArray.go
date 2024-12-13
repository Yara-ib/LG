package main

import "fmt"

func main() {
	fmt.Println("Check if num in list or not: ")
	fmt.Println(inArray([]int{5, 6, 9, 10}, 1))
	fmt.Println(inArray([]int{15, 45, 100}, 15))
	fmt.Println(inArray([]int{5000, 555, 11111, 9, 200}, 13))
	fmt.Println(inArray([]int{3000, 400, 13, 9}, 13))
}

func inArray(list []int, num int) bool {
	for _, value := range list {
		if num == value {
			return true
		}
	}
	return false
}
