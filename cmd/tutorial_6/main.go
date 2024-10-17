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

type electricEngine struct {
	kmpkwh uint8
	kwh    uint8
}

func (e electricEngine) kmsLeft() uint8 {
	return e.kmpkwh * e.kwh
}

type engine interface {
	kmsLeft() uint8
}

func canMakeIt(e engine, kms uint8) {

	if e.kmsLeft() >= kms {
		fmt.Println("Вы доедете")
	} else {
		fmt.Println("Вам не хватает топлива")
	}
}

func main() {
	var myGasEngine gasEngine = gasEngine{10, 5, owner{name: "all"}} //struct literal syntax
	myGasEngine.kmpl = 12
	canMakeIt(myGasEngine, 60)

	myElectricEngine := electricEngine{20, 5}
	canMakeIt(myElectricEngine, 101)

	// anonymus struct without defining type:
	var myEngine2 = struct {
		kmpl  uint8
		liter uint8
	}{15, 7}
	fmt.Println(myEngine2)

}
