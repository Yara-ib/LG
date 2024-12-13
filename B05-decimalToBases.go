package main

import "fmt"

func main() {
	fmt.Println(convDecimalToBase(13, 2))
	fmt.Println(convDecimalToBase(50, 2))
	fmt.Println(convDecimalToBase(50, 3))
	fmt.Println(convDecimalToBase(50, 8))
	fmt.Println(convDecimalToBase(50, 16))
	fmt.Println(convDecimalToBase(5000, 2))
	fmt.Println(convDecimalToBase(5000, 16))
	fmt.Println(convDecimalToBase(50006, 16))
	fmt.Println(convDecimalToBase(506, 16))
	fmt.Println(convDecimalToBase(6666544547, 62))
}

func convDecimalToBase(num, base int) string {
	var str string
	for num > 0 {
		// Till base 16 only >>  const charset = "0123456789ABCDEF"
		// For larger Bases | Add Extra Chars
		const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
		remainder := num % base
		str = string(charset[remainder]) + str
		num /= base
	}
	return str
}
