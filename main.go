package main

import (
	"fmt"
	"time"
)

// Structs in Go
/*
	! In the Golang ecosystem, Structs are central. They provide a mechanism to group together variables of different data types, essentially allowing the creation of custom types. Their primary function is to define and represent complex data structures in the Go language.
*/

// Animal struct will hold additional fields like Name and Age
// Structs have "fields", which can be named or anonymous
// ? Can a field be a method?
type Animal struct {
	Name string
	Age int
}

// Making a method for a struct.
	// Dynamically change the behavior of our `Animal` struct based on the animal's `name`
		// `a` is a reference within the method to the struct.
func (a Animal) Speak() string {
	switch a.Name {
	case "Dog":
		return "Woof!" // If the animal is a Dog
	case "Cat":
		return "Mew mew!" // If the animal is a Cat
	default:
		return "Unknown Animal!" // Default case for others
	}
}

// Interfaces in Go
/*
	! Interfaces provide a way to define the behavior that types must fulfill, acting as a contract or set of methods that a type should implement.
*/

// Mover interface, will have a single method called Move()
type Mover interface {
	Move() string
}

// Implementing the Mover interface in Animal Struct
// This allows instances of animals to be treated as movers.
func (a Animal) Move() string {
	return a.Name + " is moving!"
}

// Demonstraing Interface usage with a Slice of Interfaces
	// How you can use interfaces to handle a collection of differet types uniformly.
func processAnimals(animals []Mover) {
	for _, animal := range animals {
		fmt.Println(animal.Move())
	}
}

// Demonstrating buffered channels and select statement
// ! A goroutine is a function that is capable of running concurrently with other functions
func demonstrateConcurrency() {
	// * Channels provide a way for two goroutines to communicate with one another and synchronize their execution.

	c := make(chan string, 1) // * Create a buffered channel. The capacity is the number of values it can hold before becoming full, afterwhich it is blocked.

	// ! Goroutine will wait a second and then send a message to the channel.
		// Anonymous function in Golang, which is executed at the place where it is defined.
	go func() {
		time.Sleep(1 * time.Second)
		c <- "message from goroutine"
	}()

	// * Select is kind of like a switch, except instead of matching a value it will wait for a message to be recieved. The cases act like listeners for channels and then do some behavior once the messge is recieved. The message is stored in a variable.
	select {
	case msg := <-c:
		fmt.Println("Received:", msg)
	case time := <-time.After(2 * time.Second):
		// If there is no message recieved after 2 seconds, the time
		fmt.Println("Timeout: no message recieved", time)
	}
}

// Error handling
func divideNumber(x, y float64) (float64, error) {
	if y == 0.0 {
		return 0, fmt.Errorf("divide error: cannot divide by zero")
	}
	return x/y, nil

}

func main() {
	fmt.Println()

	// Demonstrating Structs and Methods
	// Use `Animal` struct to make a dog
	dog := Animal{Name: "Dog", Age: 7}
	fmt.Println("My Dog:", dog.Name, dog.Speak())

	// Use `Animal` struct to make a cat
	cat := Animal{Name: "Cat", Age: 6}
	fmt.Println("My Cat:", cat.Name, cat.Speak())

	// Demonstrating interfaces and real-world use case
	// Created a slice of Mover, which dog and cat both satisfy
	animals := []Mover{dog, cat}
	processAnimals(animals)

	// Demonstrate Concurrency
	fmt.Println("\nDemonstating Concurrency with GoRoutines and Channels")
	demonstrateConcurrency()

	// Demstonstrate Error Handling
	fmt.Println("\nDemonstrating error handling")
	result, err := divideNumber(10, 0)
	// if `err` is not `nil`, there is an error.
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Result of division:", result)

	fmt.Println()
}
