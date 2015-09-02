package main

import (
	".." // "gopkg.in/jmervine/envcfg.v1"
	"fmt"
)

func init() {
}

func main() {
	var err error
	err = envcfg.Load("_example/example.env")
	if err != nil {
		// work in _example
		err = envcfg.Load("example.env")
		if err != nil {
			panic(err)
		}
	}

	envcfg.PanicOnRequire = true

	d, _ := envcfg.Require("DATABASE_URL")
	var (
		dburl   = *d
		ignored = *(envcfg.GetOrSetBool("IGNORED", true))
		debug   = *(envcfg.GetBool("DEBUG"))
		addr    = *(envcfg.GetString("ADDR"))
		port    = *(envcfg.GetOrSetInt("PORT", 3000))
	)

	fmt.Printf("dburl   ::: %s\n", dburl)
	fmt.Printf("ignored ::: %v\n", ignored)
	fmt.Printf("debug   ::: %v\n", debug)
	fmt.Printf("addr    ::: %s\n", addr)
	fmt.Printf("port    ::: %d\n", port)
}
