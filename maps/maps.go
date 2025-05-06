package maps

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}

//? Create Maps Using the make() Function
// var a = make(map[KeyType]ValueType)

// Declare multiple key:value in map

var m = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

// var m map[string]Vertex

func LearMaps() {
	// m = make(map[string]Vertex)
	// m["Bell Labs"] = Vertex{
	// 	40.68433, -74.39967,
	// }
	// fmt.Println(m["Bell Labs"])

	// m["Bell"] = "Heo"
	// fmt.Println(m["Bell"])

	//* remove key from map
	delete(m, "Google")
	v, ok := m["Google"] // checking key exist in map or not
	fmt.Println("The value:", v, "Present?", ok)

	fmt.Println(m)
}
