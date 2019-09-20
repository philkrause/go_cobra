package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "hello",
		Short:        "hello short",
		SilenceUsage: true,
	}

	cmd.AddCommand(printTime())

	if err := cmd.Execute(); err != nil {
		println("err")
		os.Exit(1)
	}
}
