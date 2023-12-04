package main

import (
	"fmt"
	"sync"
	"time"
)

// Go Routine EXamples

// simpleMouseRoutine use WaitGroup from the sync package to handle multiple goroutines.
// A WaitGroup waits for a collection of goroutines to finish executing
func simpleMouseRoutine() {
	/*
		!Waitgroup will hold a counter that keeps track of goroutines. goroutines are tracked by adding wg.Add(1) just before they're ran.
		
		At the end of a goroutine, call wg.Done() do decrement the counter.
		
		wg.Wait() will wait for all goroutines to finish. If the counter is 0 then it will unblock the code after it.
	*/
	var wg sync.WaitGroup
	tasks := []string{"painting", "washing dishes", "doing laundry"}

	// Iterate over the tasks and create a goroutine for each task
	for _, task := range tasks {
		wg.Add(1) // Before the goroutine happens, each loop, add to the counter. This will make sure all the routines happen.
		go func(incTask string){
			defer wg.Done() // Decrementing the counter when the goroutine completes
			fmt.Printf("The Mouse is %s \n", incTask)
		}(task)
	}

	wg.Wait() // Will wait for all goroutines to finish
	fmt.Println("Simple Mouse Routine: All Tasks Completed.")
}

// Function to demonstrate the use of the select statement with channels.
	// This will include time-out handling and non-blocking channel operations.
func advancedMouseRoutine() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Start two separate goroutens simulating different tasks
	go func() {
		time.Sleep(10 * time.Second) // Simulating a shorter task
		c1 <- "Mouse is finished painting."
	}()
	
	go func() {
		time.Sleep(10 * time.Second) // Simulating a longer task
		c2 <- "Mouse is finished doing laundry."
	}()

	timeout := time.After(1 * time.Second)

	// Using select to handle multiple channel operations
	for i := 0; i < 2 ; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Advanced Mouse Routine:", msg1)
		case msg2 := <-c1:
			fmt.Println("Advanced Mouse Routine:", msg2)
		// If you dont need to use the channel's content then you do not need to assign to a variable in the case.
		case <-timeout:
			fmt.Println("Advanced Mouse Routine: Timeout Occured")
			return
		// If there is a default case, then it will be activated everytime the other cases aren't activated/
		default:
			// Default case for non-blocking select
			fmt.Println("Advanced Mouse Routine: Waiting for tasks to complete")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// Pointer Examples

// simplePointer demonstrates the use of pointers in go.
// Pointers reference the memory address of a variable.
func simplePointer() {
	cheesePieces := 10 // Mr. Chedda be like da cheese. 🧀

	// To get access address use `&` symbol in front of variable
	// To acess value at a given address use the `*` symbol in front of pointer
	pointerToCheese := &cheesePieces
	fmt.Println("Address to Mr.Chedda's cheese pieces:", pointerToCheese)
	*pointerToCheese = 20
	fmt.Printf("Simple Pointer: New number of cheese pieces = %d \n", cheesePieces)
}

// Mouse struct to represent the X,Y cordinates
type Mouse struct {
	X, Y float64
}

// Move will be a method with a pointer reciever, allowing it to modify the Mouse struct it is called on.
	// `*Mouse` is the value of the pointer for the mouse used to call this method.
func (m *Mouse) Move(dx, dy float64) {
	m.X += dx // Modifying the X coordinate to be the incoming `dx`
	m.Y += dy // Modifying the Y coordinate to be the incoming `dy`
}

// advancedPointer will keep track of mouse location.
func advancedPointer() {
	// Create a mouse from struct
	m := Mouse{X: 3, Y:4}
	m.Move(5, -2)

	// ! The %+v form shows the fields by name
	fmt.Printf("Advanced Pointer: New mouse position: %+v \n", m)

}
// Interfaces

// Interface to define the method signature for mouse activities
type MouseActivity interface {
	Perform() string // Perform method to be implemented by any type that astifies MouseActivity
}

// Struct to represent the running activity, with a speed attribute (field)
type RunActivity struct {
	Speed float64
}

// Perform method for RunActivity implements the MouseActivity interace
func (r RunActivity) Perform() string {
	// Sprintf will not print to console, will just format some string.
	return fmt.Sprintf("Running at speed %.2f", r.Speed)
}

func simpleInterface() {
	var activity MouseActivity = RunActivity{Speed: 5.0} // assign a RunActivity to a MouseACtivity interface
	fmt.Printf("Simple Interface: Mouse Activity: %s \n", activity.Perform())
}

// DigActivity represent the digging behavior, with a depth attribute
type DigActivity struct {
	Depth float64
}

// Perform Method will be bound to DigActivity structs
func (d DigActivity) Perform() string {
	return fmt.Sprintf("Digging to a depth of %.2f", d.Depth)
}

func advancedInterace() {
	// DigActivity and RunActivity both satisfy the interface MouseActivity
		// So they can both exist in a slice of MouseActivity
	activities := []MouseActivity{RunActivity{Speed: 50.0}, DigActivity{Depth: 300.0}}

	// Interating over different activities, demonstrating polymorphism
	for _, act := range activities {
		fmt.Printf("Advanced Interface: Mouse Activity: %s \n", act.Perform())
	}

}

func main() {
	// Executing all the functions to demonstrate the concepts.
	simpleMouseRoutine()
	advancedMouseRoutine()
	simplePointer()
	advancedPointer()
	simpleInterface()
	advancedInterace()
 
}
