# go-ascii

A library that provides tools for working with ASCII characters, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-ascii

[![GoDoc](https://godoc.org/github.com/reiver/go-ascii?status.svg)](https://godoc.org/github.com/reiver/go-ascii)


## Example
```
import (
	"github.com/reiver/go-ascii"
)

// ...

var byte b

// ...

switch b {
case ascii.BackSpace:
	// ...
case ascii.HorizontalTab:
	// ...
case ascii.LineFeed:

case ascii.Delete:
	// ...
}
```
