package main

import "fmt"

type gasEngine struct {
	kmpl      uint8 // 0 by default
	liter     uint8 // 0 by default
	ownerInfo owner // just owner if we don't want to add object and just add prop name string
}

type owner struct {
	name string
}

// gasEngine method
func (e gasEngine) kmsLeft() uint8 { // (e gasEngine) — определяет что это будет метод gasEngine
	return e.liter * e.kmpl // имеет доступ к полям и другим методам, как `this`
}

func main() {
	var myEngine gasEngine = gasEngine{10, 5, owner{name: "all"}} //struct literal syntax
	myEngine.kmpl = 12
	fmt.Println(myEngine.kmpl, myEngine.liter, myEngine.ownerInfo)

	// anonymus struct without defining type:
	var myEngine2 = struct {
		kmpl  uint8
		liter uint8
	}{15, 7}
	fmt.Println(myEngine2)

	fmt.Println(myEngine.kmsLeft())

}
