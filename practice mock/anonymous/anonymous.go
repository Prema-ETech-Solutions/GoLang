package anonymous

import "fmt"

func main() {
	anonymous := func() {
		fmt.Println("Anonymous function with a variable.")
	}
	anonymous()
	func() {
		fmt.Println("This is an anonymous function.")
	}()
}
