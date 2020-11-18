// Utility file to run a particular problem's file

package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {

	fmt.Println(funcMap)

	for _, arg := range os.Args[1:] {
		i, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Invalid argument '%s'\n", arg)
		} else if function, hasKey := funcMap[i]; hasKey {
			fmt.Printf("Running function 'problem%s'\n", arg)
			function()
		} else {
			fmt.Printf("Could not find function '%s'\n", arg)
		}
	}
	fmt.Println("Done.")
}

var funcMap = map[int]func(){
	1  : problem1,
	//2  : problem2,
	//3  : problem3,
	4  : problem4,
	5  : problem5,
	6  : problem6,
	//7  : problem7,
	//8  : problem8,
	9  : problem9,
	10 : problem10,
	11 : problem11,
	//12 : problem12,
	//13 : problem13,
	//14 : problem14,
	//15 : problem15,
	//16 : problem16,
	//17 : problem17,
	//18 : problem18,
	//19 : problem19,
}