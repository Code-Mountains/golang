package main 

import "fmt"

func main() {

	names := []string{} 

	names = append(names, "Keith")
	names = append(names, "Kiran", "Dangol", "Amanda")

	fmt.Println(names)
	fmt.Println("names[2] is nil: ", names[2] == "")
}

// OUTPUT 
// $ go run slices.go 
// [Keith Kiran Dangol Amanda]
// names[2] is nil:  false