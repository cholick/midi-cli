package main

import (
	"fmt"
	"os"

	"github.com/cholick/midi-cli/internal/cli"
)

func main() {
	rootCmd, err := cli.NewRootCommand()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error setting up command: %s\n", err)
		os.Exit(1)
	}

	err = rootCmd.Execute()
	if err != nil {
		// No need to print error, Cobra's SilenceErrors defaults to false ∴ already printed
		os.Exit(1)
	}
}
