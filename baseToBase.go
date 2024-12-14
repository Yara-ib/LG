package main

import "fmt"

func main() {
	fmt.Println(baseToBase("215H", 10, 2))
	fmt.Println(baseToBase("22222", 8, 16))
	fmt.Println(baseToBase("E", 16, 2))
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

func convDecimalToBase(num, base int) string {
	var str string
	for num > 0 {
		// Till base 16 only >>  const charset = "0123456789ABCDEF"
		// For larger Bases | Adding Extra Chars
		const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
		remainder := num % base
		str = string(charset[remainder]) + str
		num /= base
	}
	return str
}

func baseToBase(str string, base, convertedBase int) string {
	// Convert From base to decimal & decimal to other base
	return convDecimalToBase(basesToDec(str, base), convertedBase)
}
