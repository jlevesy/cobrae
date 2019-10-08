package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "blah",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("FAIL: %w", err)
	}
}
