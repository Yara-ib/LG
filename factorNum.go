package main

import "fmt"

func main() {
	fmt.Println(factorNum([]int{2, 3, 5}, 30))
	fmt.Println(factorNum([]int{2, 3, 5}, 28))
	fmt.Println(factorNum([]int{2, 3, 5}, 720))
	fmt.Println(factorNum([]int{3, 5}, 720))
	fmt.Println(factorNum([]int{}, 4))
}

func factorNum(primes []int, num int) []int {
	var list []int
	for i := 0; i < len(primes); i++ {
		for num%primes[i] == 0 {
			list = append(list, primes[i])
			num /= primes[i]
		}
	}
	if num > 1 {
		list = append(list, num)
	}
	return list
}
