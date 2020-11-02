// Utility file to run a particular problem's file

package main

import (
	"fmt"
	"os"
)

var funcMap = make(map[string]func())

func registerFunction(key string, function func()) {
	funcMap[key] = function
}

func main() {

	fmt.Println(funcMap)

	for _, arg := range os.Args[1:] {

		if function, hasKey := funcMap[arg]; hasKey {
			fmt.Printf("Running function 'problem%s'\n", arg)
			function()
		} else {
			fmt.Printf("Could not find function 'problem%s'\n", arg)
		}
	}
	fmt.Println("Done.")
}

func init() {
	registerFunction("1", problem1)
	registerFunction("4", problem4)
	registerFunction("5", problem5)
	registerFunction("6", problem6)
	registerFunction("9", problem9)
}
