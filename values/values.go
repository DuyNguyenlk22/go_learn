package values

import "fmt"

func Values() {
	fmt.Println("go" + "lang")
	fmt.Println(true || false)
	fmt.Println("1+1 =", 1+1)

	var a = "initial"
    fmt.Println(a)

	var b, c int = 1, 2
    fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int 
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}