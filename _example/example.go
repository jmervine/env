package main

import (
	".." // "github.com/jmervine/env"

	"fmt"
)

func init() {
	env.PanicOnRequire = true
	var err error
	err = env.Load("_example/example.env")
	if err != nil {
		// work in _example
		err = env.Load("example.env")
		if err != nil {
			panic(err)
		}
	}

	// ensure requires
	env.Require("DATABASE_URL")
}

func main() {
	fmt.Printf("dburl   ::: %s\n", env.Get("DATABASE_URL"))
	fmt.Printf("addr    ::: %s\n", env.Get("ADDR"))
	fmt.Printf("port    ::: %d\n", env.GetInt("PORT"))

	if env.GetBool("IGNORED") {
		fmt.Printf("ignored ::: %v\n", env.GetBool("IGNORED"))
	}

	if env.GetBool("DEBUG") {
		fmt.Printf("debug   ::: %v\n", env.GetBool("DEBUG"))
	}
}
