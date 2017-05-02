# random
generates fixed length random strings

## Install
`go get github.com/mj420/random`

## Usage
```golang
package main

import (
	"fmt"

	"github.com/mj420/random"
)

func main() {
	fmt.Println("Random string length of 10 is:", random.Generate(10))
}
```
