package main

import (
	"fmt"
	"parser/parser"
)

func main() {
	in := "2{a3{bc}}4{def}"
	out := parser.Parse(in)
	fmt.Println(out)
}
