# go-helpers
This is a helper library for the GoLang project.

## Features

### String helpers

| Functions	| Description	|
| - | - |
| ToLower | Converts a given string to lower case.	|
| ToUpper| Converts a given string to upper case.		|
| ToSentenceCase | Converts a given string to sentence case. 	|
| ToTitle | Converts a given string to title case or Capitalized Each Word. |
| GenerateRandomString | Generates a random string of a specified length using the alphabet characters. It uses the current time as a seed for the random number generator.	|
| GeneratePassword | Generates a random password of a specified length using a combination of lowercase letters, uppercase letters, numbers, and special characters.|
| GenerateUUID | Generates a new V7 (random) UUID. The function returns the generated UUID and an error if any.	|
| HashPassword | Generates a hashed password from a plain text password using bcrypt. The function returns the hashed password and an error if any. |
| PasswordIsValid | Checks if a given plain text password matches the hashed password. The function returns true if the passwords match, false otherwise.|
| CheckStringValue | Checks if a given string is not empty. The function returns true if the string is not empty, false otherwise.	|
| CheckStringValueToPointer | Checks if a given string is not empty and returns a pointer to the string. The function returns a pointer to the string if the string is not empty, nil otherwise.	|
| Translate | Translates a given string from source language to target language. The function returns the translated string and an error if any.|
| IsUUID | Checks if the given string is a valid UUID. The function returns true if the string is a valid UUID, false otherwise. |
| IsEmail | Checks if the given string is a valid email. The function returns true if the string is a valid email, false otherwise. |
| IsNumber | Checks if the given string is a valid number. The function returns true if the string is a valid number, false otherwise.	|
| IsPhoneNumberId | Checks if the given string is a valid phone number ID. The function returns true if the string is a valid phone number ID, false otherwise.	|

### Number helpers

| Functions	| Description	|
| - | - |
| RoundUpFloat | Rounds up the given float64 to the given uint precision. For example, if the input is 12.345 and the precision is 2, this function will return 12.35.	|
| RoundDownFloat  Rounds down the given float64 to the given uint precision. For example, if the input is 12.345 and the precision is 2, this function will return 12.34.|
| RoundFloat | Rounds the given float64 to the given uint precision, based on the rounding mode. The rounding mode is determined by the roundingUp parameter.	|
| GenerateRandomNumber | Generates a random number within a specified length. The length parameter determines the maximum value of the generated number.		|
| GetMinMaxIntArray | Get the minimum and maximum values in an array of int. The function returns the minimum and maximum values as a tuple.		|
| CheckArrayFloat64Nil | Checks if the input array of float64 is nil or empty. If not, it returns the input array. If the input array is nil or empty, it returns an empty array.	|
| CheckFloat64Value | Checks if the input pointer to a float64 is nil or empty. If not, it returns the value of the input pointer. If the input pointer is nil, it returns 0.0.	|

### Json helpers

| Functions | Description|
| - | - |
| SafeJsonMarshal | Marshals a given data to json. If the Marshal process is failed, it will return the original Marshal result and the error.|
| JsonToStruct | UnMarshals a json string to a struct. If the UnMarshal process is failed, it will return the original Marshal result and the error.	|
| JsonToString | Marshals a struct to json string. If the Marshal process is failed, it will return the original Marshal result and the error.	|

### Net helpers

| Functions | Description	|
| - | - |
| QueryEscape| Escape a string for use in a URL query.	|
| StreamToString | Reads the entire contents of an io.Reader and returns it as a string.	|
| StreamToByte | Reads the entire contents of an io.Reader and returns it as a byte slice.	|
| GetHTTPRequestJSON | Sends an HTTP request with the given method, URL, body, and timeout, and returns the response as a byte slice. The function takes an optional set of headers to include with the request.	|
| GetHTTPRequestSkipVerify | Sends an HTTP request with the given method, URL, body, and timeout, and returns the response as a byte slice. The function takes an optional set of headers to include with the request.	|


### Time helpers
- TimeUTC+7 ( [see detail](time_utc7.go))
  
  
### Response mapper
- See [Response mapper v1](https://github.com/adamnasrudin03/go-helpers/tree/main/response-mapper/v1#structure-response-api).
- Soon next version variant response structure.

### Others

| Functions | Description	|
| - | - |
| FormatErrorValidator | Formats multiple validation error messages. It takes a slice of validator.ValidationErrors and returns a slice of strings, where each string is a formatted error message.	|
| FormatErrorValidatorSingle | Formats a single validation error message. It takes a validator.ValidationErrors and returns a formatted error message.	|
| PanicRecover | This function is used to recover from a panic. It takes a string as an argument and prints it to the console.	|	


## Installation
```go
  go get github.com/adamnasrudin03/go-helpers@latest
```

## Usage
Check in go playground: https://go.dev/play/p/Qidzj-zwSa1
```go
package main

import (
	"fmt"

	help "github.com/adamnasrudin03/go-helpers"
)

func main() {
	fmt.Println(help.ToLower("Lorem ipsum dolor sit amet."))        // output; lorem ipsum dolor sit amet.
	fmt.Println(help.ToUpper("Lorem ipsum dolor sit amet."))        // output; LOREM IPSUM DOLOR SIT AMET.
	fmt.Println(help.ToSentenceCase("LOREM IPSUM DOLOR SIT AMET.")) // output; Lorem ipsum dolor sit amet.
	fmt.Println(help.ToTitle("Lorem ipsum dolor sit amet."))        // output; Lorem Ipsum Dolor Sit Amet.
}

```
other examples: https://github.com/adamnasrudin03/go-helpers/tree/main/examples


## How to Contribute
Make a pull request...

## License
Distributed under MIT License, please see license file within the code for more details.