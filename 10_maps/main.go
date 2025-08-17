package main

import (
	"fmt"
	"maps"
)

func main() {
	// Map
	var fruits = map[string]string{
		"a": "Apple",
		"b": "Banana",
		"c": "Cherry",
	}
	fmt.Println(fruits)
	fmt.Println("Length of fruits map:", len(fruits))
	fruits["d"] = "Date"
	fmt.Println(fruits)

	// accessing
	fmt.Println("Accessing fruits['b']:", fruits["b"])

	// deleting
	delete(fruits, "c")
	fmt.Println("After deleting fruits['c']:", fruits)

	// iterating
	for key, value := range fruits {
		fmt.Println("Key:", key, "Value:", value)
	}

	// make map
	var madeMap = make(map[string]int)
	madeMap["x"] = 10
	madeMap["y"] = 20
	madeMap["z"] = 30
	fmt.Println("Made map:", madeMap)

	// accessing
	fmt.Println("Accessing madeMap['y']:", madeMap["y"])

	// deleting
	delete(madeMap, "z")
	fmt.Println("After deleting madeMap['z']:", madeMap)

	// iterating
	for key, value := range madeMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	key := "x"
	value, ok := madeMap[key]
	if ok {
		fmt.Println("Accessing madeMap:", value)
	} else {
		fmt.Println("Accessing madeMap with ok:", ok)
	}

	// map equals
	anotherMap := map[string]int{
		"x": 10,
		"y": 20,
		"z": 30,
	}
	fmt.Println(maps.Equal(madeMap, anotherMap))
}
