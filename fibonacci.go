package main

import "fmt"

func main() {
	fmt.Println(fibo(11))
	fmt.Println(fibo(18))
}

func fibo(num int) ([]int, int) {
	var list []int
	var curr, numBefore, numBeforeTheBefore int

	if num <= 1 {
		list = append(list, num)
		return list, num
	}
	numBefore = 1
	numBeforeTheBefore = 0
	for i := 2; i <= num; i++ {
		curr = numBefore + numBeforeTheBefore
		list = append(list, curr)
		numBeforeTheBefore = numBefore
		numBefore = curr

	}
	return list, list[len(list)-1]
}
