// +build ignore,OMIT

package main

import "fmt"

type rect struct {
	l, b int
	name string
}

func main() {
	r := rect{l: 5, b: 3, name: "nexus"}
	fmt.Println(r)

}

func (r rect) String() string {
	return fmt.Sprintf("Name: %s\nLength: %d\nBreadth: %d\n", r.name, r.l, r.b)
}
