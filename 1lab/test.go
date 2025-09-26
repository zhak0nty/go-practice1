	package lab
	import(
		"fmt"
		"math"
	)
	func printAll(prefix string, nums ...int) {
		fmt.Println("Prefix:", prefix)
		for _, n := range nums {
			fmt.Println(n)
		}
	}

	func main() {
		printAll("Мои числа:", 1, 2, 3)
	}