package main 

import "fmt"

func main(){
	ages := map[string]int{}

	ages["Kevin"] = 33 

	if ages["Kevin"] < 18 {
		fmt.Println("Kevin can't vote")
	} else if ages["Kevin"] < 67 {
		fmt.Println("Kevin is not of retirement age")
	} else {
		fmt.Println("Enjoy retirement Kevin!")
	}

	// Switch case implementation
	switch {
	case ages["Kevin"] < 18:
		fmt.Println("Kevin can't vote")
	case ages["Kevin"] < 67:
		fmt.Println("Kevin is not of retirement age")
	default: 
		fmt.Println("Enjoy retirement Kevin!")

}