package basics

import "fmt"

/*
	Declare function using func keyword

		Without any parameter and return nothing
		func name(){}

		Takes 1 parameter and return nothing
		func name(param type){}

		Takes 1 parameter and return value
		func name(param type) (param type){}

		Takes 1 parameter and return multiple value
		func name(param type) (param1 type, param2 type){}

		Takes multiple parameter and return multiple value
		func name(param1 type, param2 type) (param1 type, param2 type){}

*/

func AboutGo() {

	fmt.Println("Go is an open source programming language that makes it easy to build simple, reliable, and efficient software")
}

func Greetings(name string) {

	s := fmt.Sprintf("Hello, %s", name)

	AboutGo()
	println(s)

}
