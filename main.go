package main

import (
	"errors"
	"fmt"
)

func main() {
	var formatString string = "hello world"
	namaVarible1 := "hello world"
	fmt.Println(namaVarible1)
	fmt.Println(formatString)
	boolVar := true
	fmt.Printf("Type: %T and value %v \n", boolVar, boolVar)
	intVar := 5
	intVar1 := int32(6)
	intVar2 := int64(8)
	floatVar1 := float32(6.0)
	floatVar2 := float64(8.9)
	runeVar := '😊'
	bytesVar := []byte("Hello World")
	complexVar := -7 + 3i
	var myInterfaceVar interface{}
	myInterfaceVar = 5
	errVar := errors.New("errror Detected")
	fmt.Printf("Type: %T and value %v \n", myInterfaceVar, myInterfaceVar)
	fmt.Printf("Type: %T and value %v \n", errVar, errVar)
	fmt.Printf("Type: %T and value %v \n", complexVar, complexVar)
	fmt.Printf("Type: %T and value %v \n", runeVar, runeVar)
	fmt.Printf("Type: %T and value %v \n", bytesVar, bytesVar)
	fmt.Printf("Type: %T and value %v \n", intVar, intVar)
	fmt.Printf("Type: %T and value %v \n", intVar1, intVar1)
	fmt.Printf("Type: %T and value %v \n", intVar2, intVar2)
	fmt.Printf("Type: %T and value %v \n", floatVar1, floatVar1)
	fmt.Printf("Type: %T and value %v \n", floatVar2, floatVar2)
}

type MethodList interface {
	MyFunction()
	MyFunction2(int) int
}
