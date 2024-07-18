package main

import (
	"fmt"

	help "github.com/adamnasrudin03/go-helpers"
)

// exampleChangeCase is an example of changing case of a string
func exampleChangeCase() {
	fmt.Println(help.ToLower("Lorem ipsum dolor sit."))        // lorem ipsum dolor sit.
	fmt.Println(help.ToUpper("Lorem ipsum dolor sit."))        // LOREM IPSUM DOLOR SIT.
	fmt.Println(help.ToSentenceCase("LOREM IPSUM DOLOR SIT.")) // Lorem ipsum dolor sit.
	fmt.Println(help.ToTitle("Lorem ipsum dolor sit."))        // Lorem Ipsum Dolor Sit.
}

// examplePassCase is an example of generating password
func examplePassCase() {
	pass := help.GeneratePassword(8)
	fmt.Println(pass)
	hashPass, err := help.HashPassword(pass)
	if err != nil {
		fmt.Println("failed hash pass", err)
	}

	pass2 := help.GeneratePassword(8)
	if help.PasswordIsValid(hashPass, pass2) {
		fmt.Println("[1] password valid")
	}
	if help.PasswordIsValid(hashPass, pass) {
		fmt.Println("[2] password valid")
	}
}

// exampleCheckValue is an example of checking if a string is not empty
func exampleCheckValue(str string) {
	fmt.Println("str", help.CheckStringValue(&str))
	fmt.Println("str", help.CheckStringValueToPointer(str))
	fmt.Println("if str is not empty")
	str = "Hello World!"
	fmt.Println("str", help.CheckStringValue(&str))
	fmt.Println("str", help.CheckStringValueToPointer(str))
}

// exampleCheckTypeValue is an example of checking the type of a value
func exampleCheckTypeValue(str string) {
	if help.IsUUID(str) {
		fmt.Println("str is UUID")
	}

	if help.IsUUID("0190c431-cd05-7702-8ebf-193ffefbc8c9") {
		fmt.Println("0190c431-cd05-7702-8ebf-193ffefbc8c9 is UUID")
	}

	if help.IsEmail(str) {
		fmt.Println("str is email")
	}
	if help.IsEmail("hello@example.com") {
		fmt.Println("hello@example.com is email")
	}

	if help.IsNumber(str) {
		fmt.Println("str is number")
	}
	if help.IsNumber("089") {
		fmt.Println("089 is number")
	}
	exampleCheckTypePhoneID(str)
}

// exampleCheckTypePhoneID is an example of checking if a string is a phone number id
func exampleCheckTypePhoneID(str string) {
	if help.IsPhoneNumberId(str) {
		fmt.Println("str is phone number id")
	}
	if help.IsPhoneNumberId("089-1234-5678") {
		fmt.Println("089-1234-5678 is phone number id")
	}
	if help.IsPhoneNumberId("+6289-1234-5678") {
		fmt.Println("+6289-1234-5678 is phone number id")
	}
	if help.IsPhoneNumberId("6289-1234-5678") {
		fmt.Println("6289-1234-5678 is phone number id")
	}
	if help.IsPhoneNumberId("+6589-1234-5678") { // +62 country SG
		fmt.Println("+6589-1234-5678 is phone number id")
	}
}

func main() {
	exampleChangeCase()
	examplePassCase()

	fmt.Println(help.GenerateRandomString(8))
	id, err := help.GenerateUUID()
	if err != nil {
		fmt.Println("failed generated uuid v7", err)
	}
	fmt.Println(id)

	exampleCheckValue("")
	exampleCheckTypeValue("Hello World!")
}
