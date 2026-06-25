package main

import "fmt"


func main() {
    // This is a placeholder for the main function.
    // You can add your code here to execute when the program runs.
	
	var Foo = func(yield func() bool) {
    	yield()
    	yield()
    	yield()
	}

    for range Foo {
        fmt.Println(".") 
    } 

	var Bar = func(yield func(int) bool) {
		yield(1)
		yield(2)
		yield(3)
	}

	for i := range Bar {
		fmt.Println(i)
	}


	var Baz = func(yield func(int, string) bool) {
		yield(1, "one")
		yield(2, "two")
		yield(3, "three")
	}

	for i, s := range Baz {
		fmt.Println(i, s)
	}
}