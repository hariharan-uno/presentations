// +build ignore,OMIT

package main

import "fmt"

type Day int

func (d Day) String() string {
	switch d {
	case 0:
		return "Sunday"
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	default:
		return "Invalid day"
	}
}

func main() {

	a := 2
	fmt.Println(a)

	b := Day(a)
	fmt.Println(b)

	fmt.Println() //Empty line

}
