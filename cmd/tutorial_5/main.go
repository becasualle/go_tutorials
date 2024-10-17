package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"

	var indexed = myString[0]
	fmt.Println(myString)
	fmt.Printf("%v, %T\n", indexed, indexed) // 114, uint8. UTF8 ASCII.

	for i, v := range myString {
		fmt.Println(i, v) // underlying representation with array of bytes
	}

	fmt.Println(len(myString)) // 8 -> not the number of characters, but the number of bytes

	var myRuneString = []rune("résumé") // [114 233 115 117 109 233] Rune is just unicode point numbers that represent the characters
	fmt.Println(myRuneString)

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v, type= %T", myRune, myRune) // myRune = 97, type= int32

	// strings are immutable and can't be changed by index once created
	var strSlice = []string{"h", "e", "l", "l", "o"}
	var myStrCopy = ""

	// here we are creating new string every time, which is inefficient
	for _, v := range strSlice {
		myStrCopy += v
	}

	fmt.Printf("\n%v", myStrCopy)

	// efficient way to perform this operation

	var strBuilder strings.Builder

	for _, v := range strSlice {
		strBuilder.WriteString(v)
	}

	var catCopy = strBuilder.String()
	fmt.Printf("\n%v", catCopy)

}
