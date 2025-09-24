[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Build Status](https://travis-ci.org/gford1000-go/hasher.svg?branch=master)](https://travis-ci.org/gford1000-go/hasher)
[![Documentation](https://img.shields.io/badge/Documentation-GoDoc-green.svg)](https://godoc.org/github.com/gford1000-go/hasher)

Hasher | Hashing of arbitrary instances using gob encoding
==========================================================

The hasher package provides the `Hash()` function that returns a `DataHash`.

The objective here is to ensure consistency of hashing, regardless of the
`interface{}` object that is passed to the `Hash()` function. If the argument is
not of type `[]byte` then `gob` is used to convert the argument to a `[]byte` slice.

The default hashing algorithm used is `sha256`, but other options can be used by passing the relevant hash type using `WithHashType()`.

An example of use is shown below:

```go
func main() {

    type Inner struct {
        B []int
    }
    type Outer struct {
        A int
        B *Inner
        C string
    }

    v := &Outer{
        A: 42,
        B: &Inner{B: []int{101, 102, 103}},
        C: "Hello",
    }

    h, _ := Hash(v) // Illustrates hashing arbitrary objects
    fmt.Println(h)  // [121 166 34 156 67 166 135 131 184 8 158 234 134 25 173 164 219 114 142 83 69 18 62 75 40 13 148 54 234 209 191 132]
}
```

Installing and building the library
===================================

This project requires Go 1.24

To use this package in your own code, install it using `go get`:

`go get github.com/gford1000-go/hasher`

Then, you can include it in your project:

`import "github.com/gford1000-go/hasher"`

Alternatively, you can clone it yourself: `git clone https://github.com/gford1000-go/hasher.git`

Testing and benchmarking
========================

To run all tests, `cd` into the directory and use:

`go test -v`
