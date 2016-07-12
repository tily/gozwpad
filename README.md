# gozwpad

Go port of [string-pad][https://github.com/youpy/string-pad) gem.

## Usage

```
package main

import (
        "fmt"
        "github.com/tily/gozwpad"
)

func main() {
        s := gozwpad.Pad("a", 140)
        fmt.Println(s)
}
```
