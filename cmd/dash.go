package cmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cpf2021-gif/mpm/cmd/dash"
	"github.com/spf13/cobra"
)

var dashCmd = &cobra.Command{
	Use:   "dash",
	Short: "View dashboard",
	Long: heredoc.Doc(`
        Display interactive dashboard.`),
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		dash.Run()
	},
}

func init() {
	rootCmd.AddCommand(dashCmd)
}
