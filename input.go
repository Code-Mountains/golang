package main 

import (
	"bufio"
	"fmt"
	"motd/message"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) 

	m := message.Greeting(name, phrase) 
	fmt.Println(m)
}

// OUTPUT 
// $ go run input.go 
// Your Greeting: Bonjour!
// Your Name: Kiran
// Bonjour!, Kiran