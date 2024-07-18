package main

import (
	"fmt"

	help "github.com/adamnasrudin03/go-helpers"
	"github.com/adamnasrudin03/go-helpers/validators"
	"github.com/go-playground/validator/v10"
)

type teamMember struct {
	FirstName      string `json:"first_name" validate:"required,min=3"`
	LastName       string `json:"last_name"`
	Email          string `json:"email" validate:"required,email"`
	UsernameGithub string `json:"username_github" validate:"required,min=3"`
}

func examplePanicRecover() {
	defer help.PanicRecover("examplePanicRecover")
	data := make([]teamMember, 0)
	fmt.Println(data[1]) // panic index out of range
}

func main() {
	examplePanicRecover()

	tm := teamMember{
		FirstName: "am",
	}
	validate := validator.New()
	err := validate.Struct(tm)
	if errors := err.(validator.ValidationErrors); errors != nil {
		temp := validators.FormatErrorValidator(errors)
		for _, v := range temp {
			fmt.Println(v)
		}
	}
}
