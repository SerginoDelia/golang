package main

// Value types
// string
// int

// format is a built-in package in Go
import "fmt"

func main() {
	// var greetingText string = "Hello world"
	greetingText := "Hello World"
	luckyNumber := 7
	otherNumebr := 10

	fmt.Println(greetingText, luckyNumber)
	fmt.Println(luckyNumber + otherNumebr)

	moreLuckyNumber := luckyNumber * 3

	fmt.Println(moreLuckyNumber)
}
