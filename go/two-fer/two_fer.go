// Package twofer implements ShareWith function
// that returns 2-fer phrase for a given name
package twofer

import "fmt"

// ShareWith returns 2-fer phrase for a given name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
