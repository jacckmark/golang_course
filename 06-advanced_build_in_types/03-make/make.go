package main

import "fmt"

func main() {
	// when we create the slices that will adjust the length based on the items number
	// dynamically we can help go by specifying how many items (approx) it will have
	// for this we use "make" function where we pass type of slice, approx length
	// (based on this go defines empty slots in the array that hold nothing for now)
	// and capacity (the capacity is the data that helps golang here in making decision
	// how much memory to allocate to this slice)
	userNames := make([]string, 2, 4)

	// when using "make" we can do things that are not possible when using regular
	// slice creation. Namely we can target the indexes (here from 0 to 2), because
	// go already allocated the memory to these spaces
	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Jerome")

	fmt.Println(userNames)

	// we can use "make" for maps as well (in maps we don't define length only
	// estimated capacity for the map) and it will also make our go application
	// a little bit more efficient
	courseRatings := make(map[string]float64, 5)
	courseRatings["go"] = 4.3
	courseRatings["angular"] = 1.2
	courseRatings["react"] = 5.0
	fmt.Println(courseRatings)

	loopOverMap(courseRatings)
	loopOverArray(userNames)
}

func loopOverMap(yourMap map[string]float64) {
	for index, value := range yourMap {
		fmt.Printf("Index: %v	Value: %v\n", index, value)
	}
}

func loopOverArray(yourArr []string) {
	for index, value := range yourArr {
		fmt.Printf("Index: %v	Value: %v\n", index, value)
	}
}
