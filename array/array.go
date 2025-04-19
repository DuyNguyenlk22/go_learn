package array

import "fmt"

func Arr() {
	myArr := [...]int{1,2,3,4,5,6}
	fmt.Printf("My Array >>> %v \n", myArr)

	mySlice := myArr[2:4]
	fmt.Printf("My slice >>> %v \n", mySlice)

	//*slice_name := make([]type, length, capacity)
	myslice1 := make([]int, 5, 10)
  	fmt.Printf("myslice1 = %v\n", myslice1)

	fmt.Printf("length = %d\n", len(myslice1))

	prices := []int{10,20,30}
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	//*slice_name = append(slice_name, element1, element2, ...)
	myslice2 := []int{1, 2, 3, 4, 5, 6}
	myslice2 = append(myslice2, 20, 21)
  	fmt.Printf("myslice2 = %v\n", myslice2)
  	fmt.Printf("length = %d\n", len(myslice2))
  	fmt.Printf("capacity = %d\n", cap(myslice2))

	//*slice3 = append(slice1, slice2...)
	myslice3 := []int{1,2,3}
 	myslice4 := []int{4,5,6}
  	myslice5 := append(myslice3, myslice4...)
  	fmt.Printf("myslice3=%v\n", myslice5)

	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice6 := arr1[1:5] // Slice array
	fmt.Printf("myslice1 = %v\n", myslice6)

	//*copy(dest, src)
	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	neededNumbers := numbers[:len(numbers)-10]
  	numbersCopy := make([]int, len(neededNumbers))
  	copy(numbersCopy, neededNumbers)

  	fmt.Printf("numbersCopy = %v\n", numbersCopy)

}