package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello")
}

func bagiAngka(x, y int) int {
	if y == 0 {
		panic("Hasil tidak dapat dibagi")
	}
	return x / y
}

func main() {
	defer func() {
		r := recover()
		fmt.Printf("Panic hasil recovery %v \n", r)
	}()
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Looping data \n")
	// }

	// j := 0
	// for j < 5 {
	// 	fmt.Printf("Looping %d\n", j)
	// 	j++
	// }

	// i := 0
	// for {
	// 	i++
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Printf("Loop %d\n", i)

	// 	if i == 10 {
	// 		break
	// 	}
	// }

	// defer hello()
	// number := []int64{1, 2, 3, 4, 5}
	// for i, v := range number {
	// 	fmt.Printf("i got index %d and value:%d \n", i, v)
	// }
	// for i := range number {
	// 	fmt.Printf("i got index %d and value:%d \n", i, number[i])
	// }

	hasilBagi := bagiAngka(5, 0)
	fmt.Println(hasilBagi)

}
