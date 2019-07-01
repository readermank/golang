package main

import "fmt"
type MyInt int

type(
	MyInt2 int
	MyInt3 int
)

var n1 MyInt = 1
var n2 MyInt = 2

var n int = 10

func main(){
	fmt.Println(n1 + n2)
}