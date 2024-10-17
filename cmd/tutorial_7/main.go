package main

import "fmt"

// POINTERS
func main() {
	var p *int32 = new(int32) // has memory adress itself and stores memory location as a value: 0x1400010200c
	var i int32

	*p = 10                                                                     // make sure it is not nil before assigning the value to the pointer
	fmt.Printf("The value p points into is: %v, at memory location: %v", *p, p) // 10, 0x1400010400c
	fmt.Printf("\nThe value of i is: %v", i)                                    // 0

	var p2 *int32 = new(int32)
	p2 = &i                                                                           // & — taking memory address, not the value
	fmt.Printf("\nThe value p points into is: %v, at memory location: %v\n", *p2, p2) // 0, 0x14000104020
	*p2 = 37                                                                          // changing the value it's points to by the memory address
	fmt.Println(i, *p2)                                                               // 37, 37

	var k int32 = 2
	k = i // copies value of i to new memory location. Values of i and k stores in separate locations
	fmt.Println(k)

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice     // just copying pointer to the array
	sliceCopy[2] = 4          // changing the value of our original slice, because under the hood slices contains pointers to underlying array
	fmt.Println(slice[2])     // 4
	fmt.Println(sliceCopy[2]) // 4

	myArr := [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of myArr is: %p\n", &myArr) // 0x1400013c000
	result := square(myArr)
	fmt.Println(result)

	result2 := squarePointer(&myArr)
	fmt.Println(result2)
}

func square(arr [5]float64) [5]float64 {
	fmt.Printf("The memory location of arr is: %p\n", &arr) // 0x1400013c030
	for i := range arr {
		arr[i] = arr[i] * arr[i]
	}

	return arr
}

// using pointer to not copy the value
func squarePointer(pArr *[5]float64) [5]float64 {
	fmt.Printf("The memory location of pArr is: %p\n", pArr) // 0x1400013c000
	for i := range pArr {
		pArr[i] = pArr[i] * pArr[i]
	}

	return *pArr
}
