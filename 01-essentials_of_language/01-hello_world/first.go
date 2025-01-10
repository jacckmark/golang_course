// package defines the space for function/functionalities (main is special because
// it is the reserved package name that tells go where execution should start)

// one module consist of many packages. To build ready to use app you need to initialize
// the module ('go mod init example.com/first-app'). This adds the .mod file and then you
// can build your app ('go build'). This in turn should create the .exe file in windows
// and other executable apps (in linux and mac)
package main

import "fmt"

// main function is also special like the main package definition. This is the starting
// point for your application. If not included your app will fail to compile. You can
// have only one main function for the whole package
func main() {
	fmt.Print("Hello world")
}

// to run the non builded application you can just type 'go run name-of-the-file.go'
// we can also run it using 'go run .' if we have done the module init (we have the
// .mod file in the project dir)
