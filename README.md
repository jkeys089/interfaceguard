# interfaceguard
interfaceguard finds subtle interface-related issues, including improper comparisons and type mismatches.

Example:
```go
package main

import "fmt"

type Greeter interface {
	SayHi() string
}

type CostcoGreeter struct {}

func (c *CostcoGreeter) SayHi() string {
	return "Welcome to Costco, I love you."
}

func main() {
	var greeter Greeter = getCachedCostcoGreeter()
	
	if greeter == nil {
		panic("greeter is nil") // we might expect this to panic, but it doesn't!
	}
	
	fmt.Println("ok")
}

func getCachedCostcoGreeter() *CostcoGreeter {
	return nil
}
```

This is a safer way to check if an interface is nil:
```go
package main

import (
	"fmt"
	"reflect"
)

type Greeter interface {
	SayHi() string
}

type CostcoGreeter struct {}

func (c *CostcoGreeter) SayHi() string {
	return "Welcome to Costco, I love you."
}

func main() {
	var greeter Greeter = getCachedCostcoGreeter()

	if isNil(greeter) {
		panic("greeter is nil") // now we catch the nil greeter and panic as expected
	}

	fmt.Println("ok")
}

func getCachedCostcoGreeter() *CostcoGreeter {
	return nil
}

func isNil(val any) bool {
	//nolint:interfaceguard
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)

	//nolint:exhaustive
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer,
		reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
```

What about comparing two interface values to each other:
```go
package main

import "fmt"

type Greeter interface {
	SayHi() string
}

type WallyWorldGreeter struct {}

func (w *WallyWorldGreeter) SayHi() string {
	return "Sorry, folks! We're closed for two weeks to clean and repair America's favorite family fun park. Sorry, uh-huh, uh-huh, uh-huh!"
}

func main() {
	var greeterA Greeter
	var greeterB Greeter = getCachedWallyWorldGreeter()
	
	if greeterA == greeterB {
		panic("greeterA == greeterB") // we might expect this to panic since both are nil, but it doesn't!
	}
	
	fmt.Println("ok")
}

func getCachedWallyWorldGreeter() *WallyWorldGreeter {
	return nil
}
```

Interface equality checks can be tricky! Here is a safer approach:
```go
package main

import (
	"fmt"
	"reflect"
)

type Greeter interface {
	SayHi() string
}

type WallyWorldGreeter struct {}

func (w *WallyWorldGreeter) SayHi() string {
	return "Sorry, folks! We're closed for two weeks to clean and repair America's favorite family fun park. Sorry, uh-huh, uh-huh, uh-huh!"
}

func main() {
	var greeterA Greeter
	var greeterB Greeter = getCachedWallyWorldGreeter()

	if isEqual(greeterA, greeterB) {
		panic("greeterA == greeterB") // now we see them as equal and panic as expected
	}
	
	fmt.Println("ok")
}

func getCachedWallyWorldGreeter() *WallyWorldGreeter {
	return nil
}

func isEqual(greeterA, greeterB Greeter) bool {
	if isNil(greeterA) && isNil(greeterB) {
		return true
	}
	
	return reflect.DeepEqual(greeterA, greeterB)
}

func isNil(val any) bool {
	//nolint:interfaceguard
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)

	//nolint:exhaustive
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer,
		reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
```

## How to use
```shell
go install github.com/jkeys089/interfaceguard/cmd/interfaceguard@latest

# default all checks enabled
interfaceguard ./...

# disable interface-to-interface comparison checks
interfaceguard -i ./...

# disable interface nil comparison checks
interfaceguard -n ./...
```