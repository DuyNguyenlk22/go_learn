package main

import (
	"example/loop"
	"example/values"
	"fmt"

	"example.com/greetings"
	"rsc.io/quote"
)


func basicTypes() {
	ToBe := false
	fmt.Printf("Type: %T value %v\n", ToBe, ToBe)
}

func main() {
    fmt.Println(quote.Opt())
	fmt.Println(quote.Hello())
	basicTypes()
	values.Values()
	loop.Loop()
	greetings.Hello("Duy Nguyen")
}