package main 

import "fmt" 

func main() {
	message := greeting("Keith", "Hello")
	fmt.Println(message)
}

// func greeting(name string, message string) string {
// 	return fmt.Sprintf("%s, %s", message, name)
// }

func greeting(name, message string) (salutation string) {
	salutation = fmt.Sprintf("%s, %s", message, name)
	return
}