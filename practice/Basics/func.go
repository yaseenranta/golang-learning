package basics

import "fmt"

func Greetings(name string) (s string) {

	s = fmt.Sprintf("Hello, %s", name)

	return
}
