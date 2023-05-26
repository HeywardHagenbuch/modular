package main

import (
	"fmt"

	"github.com/HeywardHagenbuch/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

	fmt.Println(puppy.BigBark())
	fmt.Println(puppy.BigBarks())
	puppy.From13()
	puppy.From11()
	puppy.From12()
}
