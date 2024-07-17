# go-helpers
This is a helper library for the GoLang project.

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