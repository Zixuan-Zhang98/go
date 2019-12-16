package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.firstName, p.lastName, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.firstName, sa.lastName, `says, "Shaken, not stirred."`)
}

func saySomething(h human) {
	h.speak()
}

type human interface { // interface is defined functionality that allows polymorphism
	speak()
} // if you implement interface then you are that type, since both person and secreteAgent have speak() function, they are both human

func main() {
	// primitive types:
	x := 7
	fmt.Println(x)
	fmt.Printf("%v, %T\n", x, x)
	var a int
	fmt.Println(a)

	// slice
	slice := []int{2, 4, 7, 9}
	fmt.Println(slice)

	// map
	m := map[string]int{
		"Zixuan":  21,
		"Rebecca": 21, // Need trailing comma before newline in composite literal
	}
	fmt.Println(m)

	//
	p1 := person{
		firstName: "Zixuan",
		lastName:  "Zhang",
	}

	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()

	fmt.Println("\nHere you can say something:")
	saySomething(p1)
	saySomething(sa1)
}
