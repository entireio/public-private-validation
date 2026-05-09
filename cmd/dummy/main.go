package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := newRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "dummy",
		Short: "A placeholder CLI used to demo public-code / private-checkpoint workflows",
		Long: "dummy is a no-op CLI that exists to give Entire checkpoints something to track.\n" +
			"The source code lives in a public repository while checkpoint history lives in a\n" +
			"private repository, validating the split-storage workflow documented at\n" +
			"https://docs.entire.io.",
	}
}
