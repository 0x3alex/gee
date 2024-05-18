# Gee
Evalute math and logical string expressions

## What is supported?
- Math Operations:
  - Add (`+`)
  - Substract (`-`)
  - Multiply (`*`)
  - Divide (`/`)
  - Pow (`^`)
- Logical Operations:
  - And (`&&`)
  - Or (`||`)
  - Greater (`>`)
  - Greater Than (`>=`)
  - Less (`<`)
  - Less Than (`<=`)
  - Equals (`==`)
  - Not equals (`!=`)
  - Not (`!`)
  - True Value (`True`)
  - False Value (`False`)
- Misc:
  - Braces Open (`(` )
  - Braces Close (`)`)
## What is not supported?
`precedence` is currently not supported. So multiplicate does not happen before addition, substraction, etc. 
To simulate this behaviour you can use braces.

Numbers are currently only returned as `float64`, might add the option, that `int` can be returned.

## How to use
Import the modul:
```
go get github.com/0x3alex/gee
```

And then call `gee.Eval()`.

`Eval()` returns `(int, any, error)`
- `int` describes the returned data type
  - 0 means `float64`
  - 1 means `bool`
  - 2 means `string`
- `any` the result
- `error` is set, if an error occurred while executing `eval`

## Example program
```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/0x3alex/gee"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(text)) == "exit" {
			return
		}
		t, v, err := gee.Eval(text)
		if err != nil {
			log.Fatalf(err.Error())
		}
		switch t {
		case 0:
			fmt.Printf("%.2f\n", v.(float64))
		case 1:
			println(v.(bool))
		case 2:
			println(v.(string))
		}
	}

}
```
