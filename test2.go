package main

import "fmt"

type Constructor[T any] struct {
	V0 T
}

func main() {
	var c *Constructor[string] = &Constructor[string]{V0: "hello"}
	var a any = c
	var c2 *Constructor[any] = a.(*Constructor[any])
	fmt.Println(c2.V0)
}
