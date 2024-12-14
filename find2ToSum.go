package main

import "fmt"

func main() {
	fmt.Println(find2ToSum([]int{1, 2, 3, 4, 5}, 7))
	fmt.Println(find2ToSum([]int{100, 250, 300, 4, 500}, 254))
}

// Return index of the numbers getting the total needed 
func find2ToSum(list []int, total int) (int, int) {
	for x, value := range list {
		for y, value2 := range list {
			if value == value2 {
				continue
			}
			if value+value2 == total {
				return x, y
			}
		}
	}
	return -1, -1
}
