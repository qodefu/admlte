package main

import "fmt"

type Foo struct {
	a int
}

func main() {
	sl := []Foo{Foo{1}}
	sl[0].a = 88
	fmt.Println(sl[0])

}
