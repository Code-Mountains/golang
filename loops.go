package main 

import "fmt"

func main() {
	ages := map[string]int{}
	ages["Kevin"] = 11
	ages["Keith"] = 28
	ages["James"] = 67
	ages["Michael"] = 18
	ages["Leigha"] = 16

	for name, age := range ages {
		switch age {
		case 1, 2, 3, 5, 7, 11, 13, 17, 19:
			fmt.Println(fmt.Sprintf("%s's age is a prime number!", name))
		case 16:
			fmt.Println(name, " can drive now!")
		case 18:
			fmt.Println(name, " can vote now!")			
		case 67:
			fmt.Println(name, " can retire now!")			
		default:
			fmt.Println(fmt.Sprintf("There's nothing special about %s's current age.", name))			
		}
	}


	fmt.Println("\n")

	// another implementation of for loop
	for i := 1; i <= 10; i++ {
		fmt.Println("We're counting: ", i)
	}

	fmt.Println("\n")

	// break continue implementation
	a := 0
	for a < 10 {
		fmt.Println("Count is less than 10", a)
		if a % 2 == 0 {
			a++
			fmt.Println("Increasing by +1", a)
			continue
		} else if a == 5 {
			fmt.Println("Stopping ... count is 5")
			break
		}

		fmt.Println("We're counting (again): ", a)
		a++
	}

}

// OUTPUT
// $ go run loops.go 
// Kevin's age is a prime number!
// There's nothing special about Keith's current age.
// James  can retire now!
// Michael  can vote now!
// Leigha  can drive now!


// We're counting:  1
// We're counting:  2
// We're counting:  3
// We're counting:  4
// We're counting:  5
// We're counting:  6
// We're counting:  7
// We're counting:  8
// We're counting:  9
// We're counting:  10


// Count is less than 10 0
// Increasing by +1 1
// Count is less than 10 1
// We're counting (again):  1
// Count is less than 10 2
// Increasing by +1 3
// Count is less than 10 3
// We're counting (again):  3
// Count is less than 10 4
// Increasing by +1 5
// Count is less than 10 5
// Stopping ... count is 5