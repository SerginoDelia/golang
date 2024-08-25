package main

import "fmt"

// float64 numbers with decimal places
// rune - single unicode character
// byte - single byte (ASCII character)

func main() {
	firstName := "Serge"
	lastName := "Delia"
	var currentYear = 2024
	luckyNumber := 7

	fmt.Println(firstName + ", " + lastName)
	currentYear += 1
	fmt.Println(currentYear)

	var newNumber float64 = float64(luckyNumber) / 3
	fmt.Println(newNumber)

	var smallFloat float32 = 1.3494397499347
	fmt.Println(smallFloat)

	// rune
	var firstRune rune = 'a'
	fmt.Println(firstRune)
	fmt.Print(string(firstRune) + "\n")

	// byte
	var firstByte byte = 'a'
	fmt.Println(firstByte)

	fmt.Println(firstName + " " + lastName)

	// fmt.Println("9" + 1)
}
