// ð in the name of God
// ðº S01E02 - Variables with Go language
// ð github.com/devaligh
// ð¨ð¼âð« tiwtter.com/ali_ghaffari

// ð¦
package main

// ð
import (
	"fmt"
)

// ð£
func main() {

	//  ð define variable with type
	var a int // ð¢ integer

	var ( // set many variable with one var statement
		b int     = 123 // set variable with default value
		c string        // ð¡ String type
		d float32       // 3ï¸â£âº1ï¸â£4ï¸â£ Float (32 or 64)
		e bool          // â True or flase
	)

	a = 20               // ð set value
	c, d = "Hello", 12.3 // ððset value for two vars with one line
	e = false            // âï¸ true or false

	// ð¥ show all of them in one line
	fmt.Println(a, b, c, d, e)

	// âï¸ variable naming case
	var (
		PascalCase string //ð§ð»âð
		camelCase  string //ðª
		snake_case string //ð
	)

	PascalCase = "Java developer love it â¥ï¸"
	camelCase = "PHP developer love it ð"
	snake_case = "python developer love it ð"

	fmt.Println(PascalCase)
	fmt.Println(camelCase)
	fmt.Println(snake_case)

	// ðð» define variable without define type of them
	// ð¡ and detect type of them from the value (inferred)
	inferred_var := "Hello, How Are You?"
	println(inferred_var)

	// ð CONSTANS : unchanged variable
	// ðð» Recommendation: Please use CAPITAL characters in const naming

	const PI float32 = 3.14

	const (
		SERVER_NAME string = "Test Server"
		IP          string = "127.0.01"
	)

	fmt.Println(PI * 3)
	fmt.Println("Server Name:" + SERVER_NAME)
	fmt.Println("Local IP is " + IP)

}
