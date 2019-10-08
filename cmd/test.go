package cmd

import (
	"fmt"
	"log"
	"os"

	tcmd "github.com/containous/traefik/v2/cmd"
	"github.com/containous/traefik/v2/pkg/cli"
	"github.com/containous/traefik/v2/pkg/config/env"
	"github.com/containous/traefik/v2/pkg/config/flag"
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:                "test",
		Short:              "test stuff",
		Run:                runTest,
		DisableFlagParsing: true,
	}
)

func init() {
	rootCmd.AddCommand(testCmd)
}

func isHelp(args []string) bool {
	for _, name := range args {
		if name == "--help" || name == "-help" || name == "-h" {
			return true
		}
	}
	return false
}

func toTraefikCommand(cmd *cobra.Command, cfg *tcmd.TraefikCmdConfiguration) *cli.Command {
	return &cli.Command{
		Name:          cmd.Name(),
		Description:   cmd.Short,
		Configuration: cfg,
	}
}

func runTest(cmd *cobra.Command, args []string) {
	cfg := tcmd.NewTraefikConfiguration()

	// Setting DisableFlagParsing actually disables the --help or -h detection,
	// so we need to do this by hand.
	if isHelp(args) {
		// Reusing traefik's help generation, but that's not the best.
		cli.PrintHelp(os.Stdout, toTraefikCommand(cmd, cfg))
		return
	}

	vars := env.FindPrefixedEnvVars(os.Environ(), "TEST_", cfg)

	if err := env.Decode(vars, "TEST_", &cfg.Configuration); err != nil {
		log.Fatalf("unable to decode env: %w", err)
	}

	if err := flag.Decode(args, &cfg.Configuration); err != nil {
		log.Fatalf("unable to decode flags: %w", err)
	}

	cfg.SetEffectiveConfiguration()

	fmt.Printf("config %+#v", cfg.Global)
	fmt.Printf("config %+#v", cfg.Providers.Docker)
}
