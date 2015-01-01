goregistry
==========
[![Build Status](https://travis-ci.org/josephspurrier/goregistry.svg)](https://travis-ci.org/josephspurrier/goregistry) [![Coverage Status](https://coveralls.io/repos/josephspurrier/goregistry/badge.png)](https://coveralls.io/r/josephspurrier/goregistry) [![GoDoc](https://godoc.org/github.com/josephspurrier/goregistry?status.svg)](https://godoc.org/github.com/josephspurrier/goregistry)

Golang Global Registry

Thread-safe, generic container with type conversion. Complete documentation available on [GoDoc](https://godoc.org/github.com/josephspurrier/goregistry).

Note: I would not use this package - I created it more as a learning tool for the Go language to demonstrate thread-safe operations. There are a few unconventional practices used in the GetBool() and GetInt() functions that make this package impractical. Specifically, converting between 2-3 different types for float64 types.
