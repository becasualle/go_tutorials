package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 1.9
	fmt.Println(floatNum)

	fmt.Println(float64(intNum))

	var intNum1 int = 3
	var intNum2 int = 2
	// fmt.Println(intNum1 / intNum2)                   // 1
	fmt.Println(float32(intNum1) / float32(intNum2)) // 1.5

	var myString string = "Hello \nworld"
	fmt.Println(myString)

	fmt.Println(len("A"))                    // 1 (in bytes)
	fmt.Println(len("§"))                    // 2 (in bytes)
	fmt.Println(utf8.RuneCountInString("§")) // 1 actual length

	var myRune rune = 'A'
	fmt.Println(myRune) // 65

	myText := "A" // var myText string = "A"
	myText = "B"
	fmt.Println(myText)

	int1, int2 := 1, 2 // var int1, int2 int = 1, 2
	fmt.Println(int1, int2)

	const pi float32 = 3.1415
}
