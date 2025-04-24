package structSyntax

import "fmt"

type Vertex struct {
	X int
	Y int
}

func StructSyntax() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	z := Vertex{1, 2}
	p := &z   // point to z
	p.X = 1e9 // through point p change value of z
	fmt.Printf("Value of z >>> %v", z)

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		t  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, v2, v3, t)

}
