package main

// limitations::: multi argument commands need to be run with quotes, e.g.
//
// good:
//
// $ envcfg -f .env "echo \"FOO:: \$FOO\""
// $ envcfg -f .env "ruby ./foo.rb"
// $ envcfg -f .env "go run main.go"
// $ envcfg -f .env myprog
// $ envcfg -f .env ./foo.rb
//
// bad:
//
// $ envcfg -f .env go run main.go
// $ envcfg -f .env ruby foo.rb
// $ envcfg -f .env echo foo

import (
	"github.com/jmervine/envcfg"

	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jmervine/readable"
)

var logger = readable.New().WithPrefix("[envcfg]:").WithFlags(0).WithOutput(os.Stdout)

func main() {
	var file string
	var overload bool
	var verbose bool

	flag.StringVar(&file, "f", "", "config.env file")
	flag.BoolVar(&overload, "o", false, "overload environment with confg.env")
	flag.BoolVar(&verbose, "v", false, "be verbose")
	flag.Parse()

	logger.SetDebug(verbose)

	args := flag.Args()

	if file != "" {
		if overload {
			envcfg.Overload(file)
		} else {
			envcfg.Load(file)
		}
	}

	// default to print env
	if len(args) == 0 {
		args = []string{"env"}
	}

	command := []string{"sh", "-c"}
	command = append(command, args...)

	logger.Debug(fmt.Sprintf("Running %q with %q\n", strings.Join(args, " "), file))

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Debug("Done!")
}
