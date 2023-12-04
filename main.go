package main

import (
    "fmt"
    "sync"
    "time"
)

// Concurrency Examples

// simpleMouseRoutine uses WaitGroup from the sync package to handle multiple goroutines.
// A WaitGroup waits for a collection of goroutines to finish executing.
func simpleMouseRoutine() {
    var wg sync.WaitGroup
    tasks := []string{"running", "digging", "jumping"}

    // Iterating over different tasks and creating a goroutine for each.
    for _, task := range tasks {
        wg.Add(1) // Incrementing the WaitGroup counter.
        go func(t string) {
            defer wg.Done() // Decrementing the counter when the goroutine completes.
            fmt.Printf("Mouse is %s\\n", t)
            time.Sleep(1 * time.Second) // Simulating a task with sleep.
        }(task)
    }

    wg.Wait() // Waiting for all goroutines to finish.
    fmt.Println("Simple Mouse Routine: All tasks completed.")
}

// advancedMouseRoutine demonstrates the use of the select statement with channels.
// This includes time-out handling and non-blocking channel operations.
func advancedMouseRoutine() {
    c1 := make(chan string)
    c2 := make(chan string)

    // Starting two separate goroutines simulating different tasks.
    go func() {
        time.Sleep(2 * time.Second) // Simulating a longer task.
        c1 <- "Mouse finished complex running"
    }()

    go func() {
        time.Sleep(1 * time.Second) // Simulating a shorter task.
        c2 <- "Mouse finished complex digging"
    }()

    timeout := time.After(3 * time.Second) // Setting a timeout for select.

    // Using select to handle multiple channel operations.
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("Advanced Mouse Routine:", msg1)
        case msg2 := <-c2:
            fmt.Println("Advanced Mouse Routine:", msg2)
        case <-timeout:
            fmt.Println("Advanced Mouse Routine: Timeout occurred")
            return
        default:
            // Default case for non-blocking select.
            fmt.Println("Advanced Mouse Routine: Waiting for tasks to complete")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

// Pointer Examples

// simplePointer demonstrates the use of pointers in Go.
// Pointers reference the memory address of a variable.
func simplePointer() {
    cheesePieces := 10
    pointerToCheese := &cheesePieces // Taking the address of cheesePieces.
    *pointerToCheese = 20            // Modifying the value at the pointer address.
    fmt.Printf("Simple Pointer: New number of cheese pieces = %d\\n", cheesePieces)
}

// Mouse struct represents a mouse with X and Y coordinates.
type Mouse struct {
    X, Y float64
}

// Move is a method with a pointer receiver, allowing it to modify the Mouse struct it's called on.
func (m *Mouse) Move(dx, dy float64) {
    m.X += dx // Modifying X coordinate.
    m.Y += dy // Modifying Y coordinate.
}

func advancedPointer() {
    m := Mouse{X: 3, Y: 4}
    m.Move(5, -2) // Moving the mouse by changing its coordinates.
    fmt.Printf("Advanced Pointer: New mouse position: %+v\\n", m)
}

// Interface Examples

// MouseActivity interface defines a method signature for mouse activities.
type MouseActivity interface {
    Perform() string // Perform method to be implemented by any type that satisfies MouseActivity.
}

// RunActivity struct represents a running activity with a speed attribute.
type RunActivity struct {
    Speed float64
}

// Perform method for RunActivity implements the MouseActivity interface.
func (r RunActivity) Perform() string {
    return fmt.Sprintf("Running at speed %.2f", r.Speed)
}

func simpleInterface() {
    var activity MouseActivity = RunActivity{Speed: 3.5} // Assigning a RunActivity to a MouseActivity interface.
    fmt.Printf("Simple Interface: Mouse activity: %s\\n", activity.Perform())
}

// DigActivity struct represents a digging activity with a depth attribute.
type DigActivity struct {
    Depth float64
}

// Perform method for DigActivity implements the MouseActivity interface.
func (d DigActivity) Perform() string {
    return fmt.Sprintf("Digging to depth %.2f", d.Depth)
}

func advancedInterface() {
    activities := []MouseActivity{RunActivity{Speed: 2.0}, DigActivity{Depth: 1.5}}
    // Iterating over different activities, demonstrating polymorphism.
    for _, act := range activities {
        fmt.Printf("Advanced Interface: Mouse activity: %s\\n", act.Perform())
    }
}

func main() {
    // Executing all the functions to demonstrate the concepts.
    simpleMouseRoutine()
    advancedMouseRoutine()
    simplePointer()
    advancedPointer()
    simpleInterface()
    advancedInterface()
}
