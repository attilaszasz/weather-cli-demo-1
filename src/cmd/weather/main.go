package main

import (
	"os"

	"weather-cli-demo-1/src/internal/contract"
)

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		os.Exit(contract.ExitCode(err))
	}
}
