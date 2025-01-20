package main

import "fmt"

func main() {
	// here we are defining a map. We must tell go what is the type of keys and values
	// (here they both are strings) and then add some values (key and value) in curly
	// brackets
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)

	// to access one element value in map we need to provide the key
	fmt.Println(websites["Google"])

	// we can always add new items in map, unlike in the arrays we don't define the
	// map length
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	// to remove item in map we use the "delete" function
	delete(websites, "Google")
	fmt.Println(websites)
}
