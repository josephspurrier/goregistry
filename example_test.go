// Copyright 2014 Joseph Spurrier
// Author: Joseph Spurrier (http://josephspurrier.com)
// License: http://www.apache.org/licenses/LICENSE-2.0.html

package goregistry_test

import (
	"fmt"
	"github.com/josephspurrier/goregistry"
)

// Store a value
func Example() {
	goregistry.Set("foo", "bar")
	logic()
}

// Retrieve the value
func logic() {
	if v, ok := goregistry.Get("foo"); ok {
		fmt.Println(v)
	} else {
		fmt.Println("Value not found")
	}
}
