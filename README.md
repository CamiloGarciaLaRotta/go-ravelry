# go-ravelry

![Coverage](https://img.shields.io/badge/Coverage-95.1%25-brightgreen)
[![Go Reference](https://pkg.go.dev/badge/github.com/CamiloGarciaLaRotta/go-ravelry.svg)](https://pkg.go.dev/github.com/CamiloGarciaLaRotta/go-ravelry@v0.1.0/ravelry)

Unofficial Go SDK for the [Ravelry API](https://www.ravelry.com/api)

### TL;DR

```go
package main

import (
    "fmt"

    "github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

func main() {
	// 1. setup the auth layer
	auth, err := ravelry.NewAuthFromEnv()
	if err != nil {
		panic(err)
	}

	// 2. setup the network layer
	api := ravelry.NewAPI(auth, "")

	// 3. create the API wrapper
	ravelry := ravelry.New(api, auth)

	// 4. profit
	colors, err := ravelry.ColorFamilies()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", colors)
}
```
