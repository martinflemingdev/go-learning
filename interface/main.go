package main

import (
	"fmt"
)

// declare the interface
type Greeter interface {
    LanguageName() string
    Greet(name string) string
}


// create a new struct
type GermanGreeter struct {}

// implement the interface by defining both methods
func (GermanGreeter) LanguageName() string {
	return "German"
}

func (GermanGreeter) Greet(name string) string {
	greeting := "Hallo " + name
	return greeting
}


// this is a function that takes the interface as an argument
func SayHello(Name string, greeter Greeter) string {
    hello := "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(Name)
	return hello
}


// call SayHello with implementation of the interface
func main() {
	fmt.Println(SayHello("Martin", GermanGreeter{}))
}
