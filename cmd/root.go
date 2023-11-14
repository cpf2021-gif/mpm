package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	version = "0.0.1"
)

var rootCmd = &cobra.Command{
	Use:     "mpm <command>",
	Short:   "mpm CLI",
	Long:    "Command line tool to manage your password",
	Version: version,

	SilenceUsage:  true,
	SilenceErrors: true,
}

var versionOutput = fmt.Sprintf("v%s\n", version)

var versionCmd = &cobra.Command{
	Use:  "version",
	Long: "View mpm version",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(versionOutput)
	},
}

func Execte() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
