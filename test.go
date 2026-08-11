package main

import "fmt"

type Constructor struct {
	V0 int
}

func main() {
	var c *Constructor = &Constructor{V0: 42}
	var a any = c
	var c2 *Constructor = a.(*Constructor)
	fmt.Println(c2.V0)
}
