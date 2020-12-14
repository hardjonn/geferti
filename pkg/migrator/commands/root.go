package commands

import (
	"github.com/hardjonn/geferti/migrations"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "migration",
	Short: "database migrations tool",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var helpFlag = false

// Execute executes the root command.
func Execute(s migrations.Service) error {
	// disable the "h" flag in order to avoid conflicts
	rootCmd.PersistentFlags().BoolVarP(&helpFlag, "help", "", false, "Help default flag")

	up := upCmd(s)
	down := downCmd(s)

	up.Flags().IntP("step", "s", 0, "Number of migrations to execute")
	down.Flags().IntP("step", "s", 0, "Number of migrations to rollback")

	// viper.BindPFlag("step", up.Flags().Lookup("step"))

	rootCmd.AddCommand(
		makeCmd(s),
		statusCmd(s),
		up,
		down,
	)

	return rootCmd.Execute()
}
