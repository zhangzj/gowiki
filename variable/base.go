package main

import "fmt"

//var variableName type
var oneNum int

var vname1, vname2,vname3 int

//const constantName = value
//如果需要，也可以明确指定常量的类型：
const Pi float32 = 3.1415926

var slice = []int{1,2,3}

func main() {
	slice = append(slice, int(1))

	fmt.Println(slice)

	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3

	for _, v := range numbers {
		fmt.Println(v)

	}
}


