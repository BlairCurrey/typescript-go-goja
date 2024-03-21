package main

import (
	"fmt"
	"os"

	"github.com/dop251/goja"
)

// goja minimal example
// func main () {
// 	vm := goja.New()
// 	v, err := vm.RunString("2 + 2")
// 	if err != nil {
// 			panic(err)
// 	}
// 	fmt.Println(v.Export())
// 	if num := v.Export().(int64); num != 4 {
// 			panic(num)
// 	}
// }

func main() {
	// Read the contents of the JavaScript bundle
	bundleBytes, err := os.ReadFile("../ts-module/dist/bundle.js")
	if err != nil {
			panic(err)
	}

	// Create a new JavaScript runtime
	vm := goja.New()

	// Execute the JavaScript bundle
	_, err = vm.RunScript("bundle.js", string(bundleBytes))
	if err != nil {
			panic(err)
	}

	// Call a function defined in the JavaScript bundle
	result, err := vm.RunString("MyModule.greet(\"blair\")")
	if err != nil {
			panic(err)
	}

	// Print the result
	fmt.Println(result)
}