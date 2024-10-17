package main

import "fmt"

func main() {
	var intArr [3]int32

	intArr[1] = 555

	fmt.Println(intArr[0])   // 0
	fmt.Println(intArr[1:3]) // [555, 0]

	fmt.Println(&intArr[0]) // 0x14000104020
	fmt.Println(&intArr[1]) // 0x14000104024 increments +4 each element memory placement

	anotherIntArr := [...]int32{1, 2, 3} // [1,2,3]
	// var anotherIntArr [3]int32 = [3]int32{1, 2, 3} // [1,2,3]
	fmt.Println(anotherIntArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity of %v", len(intSlice), cap(intSlice)) // The length is 3 with capacity of 3
	// if underlying array has enough room for 4 values. If not, a new array is made with enough capacity (6) and the values are copied there [4,5,6,7]
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)                                                              // [4 5 6 7]
	fmt.Printf("\nThe length is %v with capacity of %v", len(intSlice), cap(intSlice)) // The length is 4 with capacity of 6

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice) // [4 5 6 7 8 9]

	var intSlice3 []int32 = make([]int32, 3, 8)                                            // лучше заранее указывать capacity, так как операция append будет выполняться почти в три раза быстрее: without preallocation / with preallocation
	fmt.Printf("\nThe length is %v with capacity of %v\n", len(intSlice3), cap(intSlice3)) // The length is 3 with capacity of 8

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 map[string]uint8 = map[string]uint8{"Mark": 33, "Sveta": 31}
	fmt.Println(myMap2)

	var age, ok = myMap2["Alex"]
	fmt.Println(age, ok) // default value: 0, isValueExist: false

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, age: %v\n ", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n ", i, v)
	}

	// var i uint8 = 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for {
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
