package greetings

import "fmt"

func Bye(name string) string {
    message := fmt.Sprintf("CONFLICT?, %v. Welcome!", name)
    return message
}