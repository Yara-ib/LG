package main

import "fmt"

func main() {
	fmt.Println(revString("New"))
	fmt.Println(revString("Welcome Home!"))
}

func revString(text string) string {
	var revText string
	for _, letter := range text {
		revText = string(letter) + revText
	}
	return revText
}
