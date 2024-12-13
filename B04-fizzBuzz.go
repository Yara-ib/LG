package main

import "fmt"

func main() {
	fizzBuzz(5)
	fizzSwitchBuzz(5)

	fizzBuzz(50)
	fizzSwitchBuzz(50)

	fizzBuzz(15)
	fizzSwitchBuzz(15)
}

func fizzBuzz(num int) {
	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Print("Fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i)
		}
		if i < num {
			fmt.Print(", ")
		}
	}
	fmt.Println("\n")
}

func fizzSwitchBuzz(num int) {
	for i := 1; i <= num; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("Fizz Buzz")
		case i%5 == 0:
			fmt.Print("Buzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		default:
			fmt.Print(i)
		}
		if i < num {
			fmt.Print(", ")
		}
	}
	fmt.Println("\n")
}
