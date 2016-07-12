# gozwpad

[![Build Status](https://travis-ci.org/tily/gozwpad.svg?branch=master)](https://travis-ci.org/tily/gozwpad)
[![Code Climate](https://codeclimate.com/github/tily/gozwpad/badges/gpa.svg)](https://codeclimate.com/github/tily/gozwpad)
[![Issue Count](https://codeclimate.com/github/tily/gozwpad/badges/issue_count.svg)](https://codeclimate.com/github/tily/gozwpad)
[![Coverage Status](https://coveralls.io/repos/github/tily/gozwpad/badge.svg?branch=master)](https://coveralls.io/github/tily/gozwpad?branch=master)
[![GoDoc](https://godoc.org/github.com/tily/gozwpad?status.svg)](http://godoc.org/github.com/tily/gozwpad)

Go port of [string-pad](https://github.com/youpy/string-pad) gem. Pad text with random zero-width characters, mainly for posting redudant tweets.

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
