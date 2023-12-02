package main

import "testing"

//
// - File Naming: Test files in Go are named after the file they test with a '_test' suffix. 
//   For example, tests for 'main.go' should be in 'main_test.go'.
//
// - Test Function Naming: Test functions must start with 'Test' followed by a name that describes what is being tested. 
//   The function should take a single argument, t *testing.T.
//
// - Running Tests: To run tests, use the 'go test' command in the terminal. 
//   This command will run all tests in the current package.
//
// - Writing Tests: Tests typically use the t *testing.T object's methods like t.Errorf to report failures.

// TestSpeak tests the Speak method of the Animal struct
func TestSpeak(t *testing.T) {
	dog := Animal{Name: "Dog", Age: 5}
	cat := Animal{Name: "Cat", Age: 3}

	dogSpeak := dog.Speak()
	// Test to confirm that the correct output from the Speak() method occurs
	if dogSpeak != "Woof!" {
		t.Errorf("Expected Dog to say 'Woof!', got %s", dogSpeak)
	}

	catSpeak := cat.Speak()
	// Test to confirm that the correct output from the Speak() method occurs
	if catSpeak != "Mew mew!" {
		t.Errorf("Expected cat to say 'Mew mew!', got %s", catSpeak)
	}
}


// TestMove tests the Move method of the Animal struct
func TestMove(t *testing.T) {
	animal := Animal{Name: "Bird", Age: 1}
	move := animal.Move()
	expected := "Bird is moving!"

	if move != expected {
		t.Errorf("Expected '%s', got '%s'", expected, move)
	}
}