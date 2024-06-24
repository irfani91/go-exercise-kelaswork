package main

import (
	"fmt"
)

func greet(name string) (test string) {
	test = "halo " + name
	return
}

func add(x, y int) int {
	return x + y
}

// dapat dipanggil diluar package nyakaraena kapital
func AddSub(x, y int) (addRess int, subRest int) {
	addRess = x + y
	subRest = x - y
	return
}
func main() {
	fmt.Println(greet("hallo"))
	fmt.Println(add(2, 4))
	fmt.Println(AddSub(2, 5))
}
