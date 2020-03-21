package lib

import (
	"fmt"
	"strings"
)

// Greetings takes a name as single argument and return a greeting for that name
func Greetings(name string) string {
	msg := fmt.Sprintf("Hello %s!", strings.Title(name))
	return msg
}
