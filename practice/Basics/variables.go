package basics

import "fmt"

// Declare variable in go using var keyword and variable type (string|int|bool|struct|float)
// := is shorthand to declare and assign value to variable
// =  is equal is use for assign/set value to variable
// any variable is define outide function is accessible in whole project
// any variable declare inside function is only available for function

// Global scope
var (
	Name string
	Age  int
)

// string value return function declaration
func Main() string {

	Name = "GoLang"
	Age = 12

	//Available inside function only
	question := fmt.Sprintf("how many years old %v?", Name)
	answer := fmt.Sprintf("%v has %v years old", Name, Age)

	return fmt.Sprintf("%v\n%v", question, answer)
}
