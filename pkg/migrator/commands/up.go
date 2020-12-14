package commands

import (
	"fmt"

	"github.com/hardjonn/geferti/migrations"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func upCmd(s migrations.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "run up migrations",
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("PRE RUN")
			viper.BindPFlag("step", cmd.Flags().Lookup("step"))
		},
		Run: func(cmd *cobra.Command, args []string) {

			step := viper.GetInt("step")
			env := viper.GetString("env")

			fmt.Println(step)
			fmt.Println(env)

			err := s.Up(step)

			if err != nil {

			}

			// db := app.NewDB()

			// migrator, err := migrations.Init(db)
			// if err != nil {
			// 	fmt.Println("Unable to fetch migrator")
			// 	return
			// }

			// err = migrator.Up(step)
			// if err != nil {
			// 	fmt.Println("Unable to run `up` migrations")
			// 	return
			// }

		},
	}
}

// func init() {
// 	// Add "--name" flag to "create" command
// 	migrateCreateCmd.Flags().StringP("name", "n", "", "Name for the migration")

// 	// Add "--step" flag to both "up" and "down" command
// 	migrateUpCmd.Flags().IntP("step", "s", 0, "Number of migrations to execute")
// 	migrateDownCmd.Flags().IntP("step", "s", 0, "Number of migrations to execute")

// 	// Add "create", "up" and "down" commands to the "migrate" command
// 	migrateCmd.AddCommand(migrateUpCmd, migrateDownCmd, migrateCreateCmd, migrateStatusCmd)

// 	// Add "migrate" command to the root command
// 	rootCmd.AddCommand(migrateCmd)
// }
