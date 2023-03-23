package greetings

import "fmt"

// Hello returns the message "Hello, World!"
func Hello() string {
	return "Hello, World!"
}

// HelloToJason returns the legendary wittertainment message
func HelloToJason() string {
	return "Hello to Jason Isaacs!"
}

// HelloTo returns a greeting message to the given name
func HelloTo(name string) string {
	if len(name) == 0 {
		name = "stranger"
	}
	return fmt.Sprintf("Hello to %v!", name)
}