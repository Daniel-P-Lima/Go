package greetings 

import "fmt"

// Hello returns a greeting for the name person
func Hello(name string) string {
	// Return a greeting tha embeds the ame in a message
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message
}