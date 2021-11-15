package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Person struct {
	FirstName string `json:"first_name" validate:"required,min=3"`
	LastName  string `json:"last_name" validate:"required"`
	Age       int    `json:"age" validate:"required"`
}

type Response struct {
	Status           int
	Message          string
	Error            string
	ValidationErrors interface{}
}

func FailedResponse(status int, message string, err interface{}) (response Response) {

	response.Status = status
	response.Message = message

	switch t := err.(type) {
	case string:
		response.Error = t
		response.ValidationErrors = nil

	default:
		response.Error = ""
		response.ValidationErrors = err
	}
	return
}

func (p *Person) formatError() (validationMessages map[string]map[string]interface{}) {

	validationMessages = map[string]map[string]interface{}{
		"FirstName": {
			"required": "First Name is required",
			"min":      "First Name have minimum 3 character",
		},
		"LastName": {
			"required": "Last Name is required",
		},
		"Age": {
			"required": "Last Name is required",
		},
	}

	return
}

func main() {

	p := Person{
		FirstName: "",
		LastName:  "",
		Age:       18,
	}

	validationMessages := p.formatError()
	holdValidationError := map[string]interface{}{}

	validate := validator.New()

	err := validate.Struct(p)

	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {

			holdValidationError[e.Field()] = validationMessages[e.Field()][e.ActualTag()]

		}
		res := FailedResponse(200, "validation Failed", holdValidationError)
		//		fmt.Println(res)

		var w http.ResponseWriter
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		fmt.Println(json.NewEncoder(w).Encode(&res))

	}

}
