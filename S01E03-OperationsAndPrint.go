// 🕋 in the name of God
// 📺 S01E03 - Operations and Print methods
// 📃 github.com/devaligh
// 👨🏼‍🏫 tiwtter.com/ali_ghaffari

// 📦
package main

// 📚
import (
	"fmt"
	"strconv" // for convertion (string to int)
)

// 💣
func main() {

	//  🎚 define variables
	var (
		a int    = 20
		b int    = 10
		c string = "Hello"
		d string = "World"
	)

	// 🧮 math functions
	fmt.Printf("a + b = %d \n", a+b)
	fmt.Printf("a - b = %d \n", a-b)
	fmt.Printf("a * b = %d \n", a*b)
	fmt.Printf("a / b = %d \n", a/b)
	fmt.Printf("a mod b = %d \n", a%b) // Divide remaining

	// ➕ concat two string
	fmt.Println(c + " " + d + "!") // join Hello + SPACE + World + DANGER_SIGN

	// 🔄 convert int to string
	var test_string string = strconv.FormatInt(int64(a), 10)
	fmt.Println(test_string)

	// 🔄 convert string to int
	test_int := "100"
	int_var, err := strconv.Atoi(test_int)
	fmt.Println(int_var, err)

	// ⁉️ get value and type of value
	fmt.Printf("a is %v and type is %T", a, a)

}
