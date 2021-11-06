// 🕋 in the name of God
// 📺 S01E02 - Variables with Go language
// 📃 github.com/devaligh
// 👨🏼‍🏫 tiwtter.com/ali_ghaffari

// 📦
package main

// 📚
import (
	"fmt"
)

// 💣
func main() {

	//  🎚 define variable with type
	var a int // 🔢 integer

	var ( // set many variable with one var statement
		b int     = 123 // set variable with default value
		c string        // 🔡 String type
		d float32       // 3️⃣⏺1️⃣4️⃣ Float (32 or 64)
		e bool          // ✅ True or flase
	)

	a = 20               // 📝 set value
	c, d = "Hello", 12.3 // 📝📝set value for two vars with one line
	e = false            // ⛔️ true or false

	// 🖥 show all of them in one line
	fmt.Println(a, b, c, d, e)

	// ⁉️ variable naming case
	var (
		PascalCase string //🧑🏻‍🎓
		camelCase  string //🐪
		snake_case string //🐍
	)

	PascalCase = "Java developer love it ♥️"
	camelCase = "PHP developer love it 💙"
	snake_case = "python developer love it 💚"

	fmt.Println(PascalCase)
	fmt.Println(camelCase)
	fmt.Println(snake_case)

	// 👏🏻 define variable without define type of them
	// 💡 and detect type of them from the value (inferred)
	inferred_var := "Hello, How Are You?"
	println(inferred_var)

	// 💎 CONSTANS : unchanged variable
	// 👍🏻 Recommendation: Please use CAPITAL characters in const naming

	const PI float32 = 3.14

	const (
		SERVER_NAME string = "Test Server"
		IP          string = "127.0.01"
	)

	fmt.Println(PI * 3)
	fmt.Println("Server Name:" + SERVER_NAME)
	fmt.Println("Local IP is " + IP)

}
