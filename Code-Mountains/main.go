/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import "Code-Mountains/cmd"

func main() {
	cmd.Execute()
}


// OUTPUT 

// $ DEBUG=true ./Code-Mountains -n "Kiran" -g "Aloha"
// Name:  Kiran
// Greeting:  Aloha
// Prompt:  false

// $ DEBUG=true go run main.go 
// Usage:
//   Code-Mountains [flags]

// Flags:
//   -g, --greeting string   Greeting to use in message
//   -h, --help              help for Code-Mountains
//   -n, --name string       Name to use in message
//   -v, --preview           Preview message instead of writing /etc/motd
//   -p, --prompt            Prompt for name and greeting
// exit status 1

// $ DEBUG=true go run main.go -n "Kiran" -g "Aloha" -v
// Name:  Kiran
// Greeting:  Aloha
// Prompt:  false

