package main 

import "fmt"

func main(){
	ages := map[string]int{}

	ages["Kevin"] = 61 

	if ages["Kevin"] < 67 {
		fmt.Println("Kevin is not of retirement age")
	} else {
		fmt.Println("Kevin is of retirement age")
	}
}