# go-ascii

Package **ascii** provides tools for working with ASCII characters, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-ascii

[![GoDoc](https://godoc.org/github.com/reiver/go-ascii?status.svg)](https://godoc.org/github.com/reiver/go-ascii)


## Example

Here is an example of using the longer versions of the ASCII constants:

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
	// ...
case ascii.Delete:
	// ...
default:
	// ...
}
```

Here is an example of using the longer versions of the ASCII constants:

```
import (
	"github.com/reiver/go-ascii"
)

// ...

var byte b

// ...

switch b {
case ascii.BS
	// ...
case ascii.HT:
	// ...
case ascii.LF:
	// ...
case ascii.DEL:
	// ...
default:
	// ...
}
```

## Import

To import package **ascii** use `import` code like the following:
```
import "github.com/reiver/go-ascii"
```

## Installation

To install package **ascii** do the following:
```
GOPROXY=direct go get github.com/reiver/go-ascii
```

## Author

Package **ascii** was written by [Charles Iliya Krempeaux](http://reiver.link)
