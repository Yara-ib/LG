package main

import "fmt"

func main() {
	fmt.Println(sumList([]int{1, 5, 5, 6}))
	fmt.Println(sumList([]int{11, 55, 105, 86}))
	fmt.Println(sumListB([]int{11, 55, 105, 86}))
}

func sumList(list []int) int {
	sum := 0
	for _, val := range list {
		sum += val
	}
	return sum
}

func sumListB(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + sumList(list[1:])
}
