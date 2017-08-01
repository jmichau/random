[![Go Report Card](https://goreportcard.com/badge/github.com/mikejav/random)](https://goreportcard.com/report/github.com/mikejav/random)
# random
generates fixed length random strings. Useful e.g. for generating session ID

## Install
```shell
go get github.com/mikejav/random
```

## Usage
```golang
...
random.Generate(length)
...
```

## Example
```golang
package main

import (
	"fmt"

	"github.com/mikejav/random"
)

func main() {
	fmt.Println("Random string length of 10 is:", random.Generate(10))
}
```
