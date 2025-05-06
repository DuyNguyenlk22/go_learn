package main

import (
	"example/maps"

	"fmt"
)

func basicTypes() {
	ToBe := false
	fmt.Printf("Type: %T value %v\n", ToBe, ToBe)
}

func learnDefer() {
	//? Before func learnDefer end => execute defer
	//? LIFO
	defer fmt.Println("World")

	fmt.Println("Hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	// fmt.Println(quote.Opt())
	// fmt.Println(quote.Hello())
	//learnDefer()
	// basicTypes()
	// values.Values()
	// loop.Loop()
	// greetings.Hello("Duy Nguyen")
	// array.Arr()
	//conditions.LearnConditions()
	//conditions.LearnSwitchCase()
	// pointers.Pointers()
	// structSyntax.StructSyntax()
	// slices.LearSlices()
	maps.LearMaps()
}
