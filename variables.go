package main 

import "fmt" 

func main() {
	var myInt int = 16
	var val, ok = "yes", true

	fmt.Println("myInt is:", myInt)
	fmt.Println("myInt times two:", myInt * 2)
	fmt.Println("val is:", val)
	fmt.Println("ok is:", ok)
}

// OUTPUT 
// $ go run variables.go 
// myInt is: 16
// myInt times two: 32
// val is: yes
// ok is: true