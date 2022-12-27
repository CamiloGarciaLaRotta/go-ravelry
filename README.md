# go-ravelry

![Coverage](https://img.shields.io/badge/Coverage-97.3%25-brightgreen)
[![Go Reference](https://pkg.go.dev/badge/github.com/CamiloGarciaLaRotta/go-ravelry.svg)](https://pkg.go.dev/github.com/CamiloGarciaLaRotta/go-ravelry)

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

### Status

Contributions welcome!

- ✅ implemented
- ❌ not yet implemented

| API Endpoint           | Status |
| ---------------------- | ------ |
| color_families         | ✅     |
| current_user           | ✅     |
| search                 | ✅     |
| yarn_weights           | ✅     |
| app                    | ❌     |
| bundled_items          | ❌     |
| bundles                | ❌     |
| carts                  | ❌     |
| comments               | ❌     |
| deliveries             | ❌     |
| designers              | ❌     |
| extras                 | ❌     |
| favorites              | ❌     |
| fiber                  | ❌     |
| fiber_attribute_groups | ❌     |
| fiber_attributes       | ❌     |
| fiber_categories       | ❌     |
| forum_posts            | ❌     |
| forums                 | ❌     |
| friends                | ❌     |
| groups                 | ❌     |
| in_store_sales         | ❌     |
| library                | ❌     |
| messages               | ❌     |
| needles                | ❌     |
| packs                  | ❌     |
| pages                  | ❌     |
| pattern_attributes     | ❌     |
| pattern_categories     | ❌     |
| pattern_sources        | ❌     |
| patterns               | ❌     |
| people                 | ❌     |
| photos                 | ❌     |
| product_attachments    | ❌     |
| products               | ❌     |
| projects               | ❌     |
| queue                  | ❌     |
| saved_searches         | ✅     |
| shops                  | ❌     |
| stash                  | ❌     |
| stores                 | ❌     |
| topics                 | ❌     |
| upload                 | ❌     |
| volumes                | ❌     |
| yarn_attributes        | ✅     |
| yarn_companies         | ✅     |
| yarns                  | ❌     |
