# random
generates fixed length random strings. Useful e.g. for generating session ID

## Install
```shell
go get github.com/mj420/random
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

	"github.com/mj420/random"
)

func main() {
	fmt.Println("Random string length of 10 is:", random.Generate(10))
}
```
