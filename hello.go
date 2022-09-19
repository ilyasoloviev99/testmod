package testmod

import "fmt"

// Hi returns a friendly greeting
func Hi(name, surname string) string {
	return fmt.Sprintf("Hi, %s %s", name, surname)
}
