package main

import (
	"fmt"

	help "github.com/adamnasrudin03/go-helpers"
)

type teamMember struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	UsernameGithub string `json:"username_github"`
}

func main() {
	tm := teamMember{
		FirstName: "Adam",
		LastName:  "Nasrudin",
		Email:     "adamnasrudin@example.com",
	}

	a, err := help.SafeJsonMarshal(tm)
	if err != nil {
		fmt.Println("SafeJsonMarshal", err)
	}
	b := string(a)
	fmt.Println(b)

	c := teamMember{}
	err = help.JsonToStruct(b, &c)
	if err != nil {
		fmt.Println("JsonToStruct", err)
	}
	fmt.Println(c)

	d, err := help.JsonToString(tm)
	if err != nil {
		fmt.Println("JsonToString", err)
	}
	fmt.Println(d)

}
