package main

import "fmt"

func main() {
	fmt.Println(basesToDec("111", 16))
	fmt.Println(basesToDec("EF11", 16))
	fmt.Println(basesToDec("1111", 2))
	fmt.Println(basesToDec("22222", 8))
	fmt.Println(basesToDec("215H", 10))
}

func basesToDec(str string, base int) int {
	var convertedNum int
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	multiplier := 1
	for i := len(str) - 1; i >= 0; i-- {
		idx := -1
		for j, value := range charset {
			if value == rune(str[i]) {
				idx = j
				break
			}
		}
		convertedNum += multiplier * idx
		multiplier *= base
	}
	return convertedNum
}
