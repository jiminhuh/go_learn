package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	foo("Paris")
	fmt.Println("_________")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo(place string) {
	fmt.Println("I'm in " + place)
}
