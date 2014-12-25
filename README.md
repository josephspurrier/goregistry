goregistry
==========

Golang Global Registry

Implements the singleton pattern so the same instance can be used anywhere.

## Add the import

```go
import "github.com/josephspurrier/goregistry"
```

## Create an instance

```go
reg := goregistry.Registry()
```

## Store a value

```go
reg.Set("foo", "bar")
```

## Retrieve a value

```go
value, err := reg.Get("foo")
```

## Retrieve a string value

```go
value, err := reg.GetString("foo")
```

## Retrieve a bool value

```go
value, err := reg.GetBool("foo")
```

## Retrieve a int value

```go
value, err := reg.GetInt("foo")
```

## Full example

```go
package main

import (
	"fmt"
	"github.com/josephspurrier/goregistry"
)

func test() {
	reg := goregistry.Registry()
	if v, ok := reg.Get("foo"); ok {
		fmt.Println(v)
	} else {
		fmt.Println("Value not found")
	}
}

func main() {
	reg := goregistry.Registry()
	reg.Set("foo", "bar")

	test()
}
```