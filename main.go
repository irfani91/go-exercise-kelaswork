package main

import (
	"fmt"
)

func incrementPointer(var1 *int) {
	*var1 = *var1 + 1
}

func increment(var1 int) {
	var1 = var1 + 1
}
func main() {
	var pointerSaya *string
	var kalimat = "Hello Saya"
	pointerSaya = &kalimat

	var pointerSaya2 = new(string)

	*pointerSaya = "Belajar Golang"
	fmt.Println(pointerSaya)
	fmt.Println(*pointerSaya)
	fmt.Println(kalimat)
	fmt.Println(pointerSaya2)

	var1 := 1
	fmt.Printf("my value is %d\n", var1)
	increment(var1)
	fmt.Printf("my value is %d\n", var1)
	incrementPointer(&var1)
	fmt.Printf("my value is %d\n and address %X \n", &var1, var1)
	var pointerSaya3 *string
	//fmt.Printf("Nilai nya %v", *pointerSaya3)
	fmt.Printf("Nilai nya %v \n", pointerSaya3)
	pointerSaya3 = new(string)
	*pointerSaya3 = "sakit"
	fmt.Printf("Nilai nya %v\n", *pointerSaya3)

	fmt.Printf("Nilainya %v\n", kalimat)
	pointerSaya3 = &kalimat
	*pointerSaya3 = "ubah "
	fmt.Printf("Nilainya %v\n", kalimat)
	fmt.Printf("Nilai nya %v", *pointerSaya3)
}
