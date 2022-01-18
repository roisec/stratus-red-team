package main

import (
	_ "github.com/datadog/stratus-red-team/internal/attacktechniques"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "stratus-red-team",
}

func init() {
	listCmd := buildListCmd()
	showCmd := buildShowCmd()
	warmupCmd := buildWarmupCmd()
	detonateCmd := buildDetonateCmd()
	statusCmd := buildStatusCmd()
	cleanupCmd := buildCleanupCmd()
	versionCmd := buildVersionCmd()

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(warmupCmd)
	rootCmd.AddCommand(detonateCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(cleanupCmd)
	rootCmd.AddCommand(versionCmd)
}

func main() {
	rootCmd.Execute()
}