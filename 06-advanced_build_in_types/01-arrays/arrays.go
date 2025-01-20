package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	// here we are initializing the array. We define the length of 4 and type stored
	// there, then we pass the values we want to store there
	prices := [4]float64{10.99, 11.34, 45.00, 20.11}

	// here we define the array but don't add any values there. When printed we will
	// get an array with 4 empty slots
	var productNames [4]string

	fmt.Println(prices)
	fmt.Println(productNames)

	// we can fill in only one place in the array if we want
	productNames = [4]string{"honda"}
	fmt.Println(productNames)

	// getting one value from array using index (first element)
	fmt.Println(prices[0])

	// we can set the only one slot value using index (here setting the third element)
	productNames[2] = "kia"
	fmt.Println(productNames)

	// we can also create the slices of our arrays (parts of lists). They are very
	// efficient because they are not creating a copy of the original array. Instead
	// you have an original list intact and your slice stores the reference to the
	// items it is storing from the original array
	limitedPrices := prices[1:3]
	fmt.Println(limitedPrices)

	// slices give us also an ability to select the items without providing fist
	// or second index whatsoever (here getting all elements till index 2 exclusive
	// and getting all the items from index 2 till the end exclusive)
	fmt.Println(prices[:2])
	fmt.Println(prices[2:])

	// like in other languages we have to watch out for the indexes. We cannot pick
	// index greater than the list length and also cannot use -1

	// to check the length of your slice you use "len" function
	fmt.Println(len(limitedPrices))

	// to check the defined capacity of your array you use "cap" function (for slices
	// it will show you the capacity of underlying array). When we append to the array
	// in go the length increases until it reaches capacity, then the capacity is
	// increased
	fmt.Println(cap(limitedPrices))

	// for arrays where we don't know the length ahead of time we can create the
	// array without specifying the capacity (go will create a slice instead the
	// regular array and will handle the capacity by itself when we will be putting
	// new items there)
	monkeys := []string{"Joe", "Mary", "James"}
	fmt.Println(monkeys)
	fmt.Println(monkeys[0:1])

	// to add new items to the array in go we use the "append" (you cannot use the
	// index notation to add new items). This function does not mutate the original
	// array, it just return the new slice altogether
	newMonkeys := append(monkeys, "Robert", "Barbara")
	fmt.Println(newMonkeys)

	// to remove the items we create new slice based on old slice or if it is an
	// item that resides in the middle of the array we loop over the array and
	// omit the item we don't want to have in the new array
	newerMonkeys := monkeys[:2]
	fmt.Println(newerMonkeys)

	// to add one array to the another we cannot just use "append" and pass the var
	// as the second argument to it. We have to unpack the list
	newRecruits := []string{"Karen", "Daniel", "Sebastian"}
	evenNewerMonkeys := append(newMonkeys, newRecruits...)
	fmt.Println(evenNewerMonkeys)
}
