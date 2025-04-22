package conditions

import "fmt"

func LearnConditions() {
	x := 30
	if x >= 10 {
		fmt.Println("x is larger than or equal to 10.")
	} else if x > 20 {
		fmt.Println("x is larger than 20.")
	} else {
		fmt.Println("x is less than 10.")
	}
}

func LearnSwitchCase() {
	x := 30
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	default:
		fmt.Println("x is neither 1 nor 2")
	}
}
