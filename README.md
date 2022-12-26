# go-ravelry

![Coverage](https://img.shields.io/badge/Coverage-93.0%25-brightgreen)
[![Go Reference](https://pkg.go.dev/badge/github.com/CamiloGarciaLaRotta/go-ravelry.svg)](https://pkg.go.dev/github.com/CamiloGarciaLaRotta/go-ravelry)

Unofficial Go SDK for the [Ravelry API](https://www.ravelry.com/api)

### TL;DR
For more information on how to use the API wrapper, see [the GoDocs](https://pkg.go.dev/github.com/CamiloGarciaLaRotta/go-ravelry).

```go
package main

import (
    "fmt"

    "github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

func main() {
    api := ravelry.NewAPI(auth, "")

    // you can also use ravelry.NewAuth to directly pass the credentials
    auth, err := ravelry.NewAuthFromEnv()
    if err != nil {
        panic(err)
    }

    ravelry := ravelry.New(api, auth)

    res, err := ravelry.ColorFamilies()
    if err != nil {
        panic(err)
    }

    fmt.Printf("%v\n", res)
}
```
