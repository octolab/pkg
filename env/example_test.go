package env_test

import (
	"fmt"
	"os"

	"go.octolab.org/env"
)

func Example() {
	vars := env.From(os.Environ())
	val, present := vars.Lookup(env.Go111Module)
	if present {
		fmt.Println("env name:", val.Name())
		fmt.Println("env value:", val.Value())
		fmt.Println("string form:", val)
	}
	// output:
	// env name: GO111MODULE
	// env value: on
	// string form: GO111MODULE=on
}
