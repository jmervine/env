package main

import (
	".." // "github.com/jmervine/env"

	"fmt"
)

func init() {
}

func main() {
	var err error
	err = env.Load("_example/example.env")
	if err != nil {
		// work in _example
		err = env.Load("example.env")
		if err != nil {
			panic(err)
		}
	}

	env.PanicOnRequire = true

	d, _ := env.Require("DATABASE_URL")
	var (
		dburl   = d
		ignored = env.GetOrSetBool("IGNORED", true)
		debug   = env.GetBool("DEBUG")
		addr    = env.GetString("ADDR")
		port    = env.GetOrSetInt("PORT", 3000)
	)

	fmt.Printf("dburl   ::: %s\n", dburl)
	fmt.Printf("ignored ::: %v\n", ignored)
	fmt.Printf("debug   ::: %v\n", debug)
	fmt.Printf("addr    ::: %s\n", addr)
	fmt.Printf("port    ::: %d\n", port)
}
