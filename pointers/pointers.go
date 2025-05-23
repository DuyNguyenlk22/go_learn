package pointers

import "fmt"

var p *int // declare a pointer p

func Pointers() {
	i, j := 42, 2701
	p = &i          // point to i
	fmt.Println(*p) // read i through the pointer
	fmt.Println(p)  // memory address of p
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j       //point to j
	*p = *p / 37 //divide j through the pointer
	fmt.Println(j)
}
