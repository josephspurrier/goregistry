goregistry
==========

Golang Global Registry

Generic container with type conversion.

## Add the import

```go
import "github.com/josephspurrier/goregistry"
```

## Store a value

```go
goregistry.Set("foo", "bar")
```

## Retrieve a value

```go
value, err := goregistry.Get("foo")
```

## Retrieve a string value

```go
value, err := goregistry.GetString("foo")
```

## Retrieve a bool value

```go
value, err := goregistry.GetBool("foo")
```

## Retrieve a int value

```go
value, err := goregistry.GetInt("foo")
```

## Full example

```go
import (
	"github.com/josephspurrier/goregistry"
	"fmt"
)

func main {
	goregistry.Set("foo", "bar")

	if v, ok := goregistry.Get("foo"); ok {
		fmt.Println(v);
	} else {
		fmt.Println("Value not found");
	}
}
```