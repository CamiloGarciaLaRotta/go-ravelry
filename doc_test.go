package go_ravelry

import (
	"fmt"

	"github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"
)

func Example() {
	// import "github.com/CamiloGarciaLaRotta/go-ravelry/ravelry"

	// 1. setup the auth layer
	auth, err := ravelry.NewBasicAuthFromEnv()
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
