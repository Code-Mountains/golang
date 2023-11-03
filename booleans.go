package main 

import "fmt"

func main() {
	fmt.Println("Great than:", 1 > 3)
	fmt.Println("Less than:" , 2 < 3)
	fmt.Println("Greater than OR equal:", 2 >= 2)
	fmt.Println("Less than OR equal:", 4 <= 4)
	fmt.Println("Equivalent:", 4.0 == 4)
	fmt.Println("Not equivalent:", 4.1 != 4.1)
}

// OUTPUT 
// $ go run booleans.go 
// Great than: false
// Less than: true
// Greater than OR equal: true
// Less than OR equal: true
// Equivalent: true
// Not equivalent: false