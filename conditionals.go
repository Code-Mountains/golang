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

	switch ages["Kevin"] {
	case 1, 2, 3, 5, 7, 11, 13, 17, 19:
		fmt.Println("Kevin's age is a prime number!")
	case 16:
		fmt.Println("Kevin can drive now")
	case 18:
		fmt.Println("Kevin can vote now")
	case 26:
		fmt.Println("Kevin can retire now")		
	default:
		fmt.Println("There's nothing special about Kevin's current age")
	}
}