package slices

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// Create a 2D slice of uint8
	image := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// Try different functions like:
			// row[x] = uint8((x + y) / 2)
			// row[x] = uint8(x * y)
			row[x] = uint8(x ^ y) // using XOR for interesting pattern
		}
		image[y] = row
	}
	return image
}

func LearSlices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pic.Show(Pic)
}
