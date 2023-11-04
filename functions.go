package main 

import "fmt" 

func main() {
	message := greeting("Keith", "Hello")
	fmt.Println(message)
}

func greeting(name, message string) (salutation string) {
	salutation = fmt.Sprintf("%s, %s", message, name)
	return
}

// OUTPUT 
// $ go run functions.go 
// Hello, Keith