// 🕋 in the name of God
// 📺 S01E04 - Array and Slices by Go language
// 📃 github.com/devaligh
// 👨 tiwtter.com/ali_ghaffari

// 📦 set package name
package main

import "fmt"

// 📚 import fmt library for use Print in screen function

// 💣 main is startup function by default
func main() {

	// 📒 Array
	var array_a = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array_a)    // OUPUT: [1 2 3 4 5]
	fmt.Println(array_a[3]) // OUTPUT: 4

	// define array without type (inferred)
	var array_b = [...]string{"Iran", "US", "Germany"}
	fmt.Println(array_b[0]) // OUTPUT: Iran

	// define array with index of array
	var array_c = [4]int{1: 123, 3: 456}
	fmt.Println(array_c) // OUTPUT: [0 123 0 456]

	// lenght
	fmt.Println(len(array_c))

	// 📚 Slice
	myslice := []int{}
	myslice = append(myslice, 1)
	myslice = append(myslice, 2)
	myslice = append(myslice, 3)
	fmt.Println(myslice) // OUTPUT:		[1 2 3]

	// Modify
	myslice[2] = 4
	fmt.Println(myslice) // OUTPUT:		[1 2 4]

	// len & cap
	fmt.Println(len(myslice)) // Lenght of the slice
	fmt.Println(cap(myslice)) // Capacity of the slice

}
