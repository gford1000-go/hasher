[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Build Status](https://travis-ci.org/gford1000-go/hasher.svg?branch=master)](https://travis-ci.org/gford1000-go/hasher)
[![Documentation](https://img.shields.io/badge/Documentation-GoDoc-green.svg)](https://godoc.org/github.com/gford1000-go/hasher)


Hasher | Geoff Ford
===================

The hasher package provides the `Hash()` function that returns a `DataHash`.

An example of use is available in GoDocs.

The objective here is to ensure consistency of hashing, regardless of the
`interface{}` object that is passed to the `Hash()` function. If the argument is
not of type `[]byte` then `gob` is used to convert the argument to a `[]byte` slice.

The hashing algorithm used is `sha256.Sum256`.


Installing and building the library
===================================

This project requires Go 1.14.2

To use this package in your own code, install it using `go get`:

    go get github.com/gford1000-go/hasher

Then, you can include it in your project:

	import "github.com/gford1000-go/hasher"

Alternatively, you can clone it yourself:

    git clone https://github.com/gford1000-go/hasher.git

Testing and benchmarking
========================

To run all tests, `cd` into the directory and use:

	go test -v

