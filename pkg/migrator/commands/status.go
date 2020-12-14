package commands

import (
	"github.com/hardjonn/geferti/migrations"

	"github.com/spf13/cobra"
)

func statusCmd(s migrations.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "display status of each migrations",
		Run: func(cmd *cobra.Command, args []string) {
			// db := app.NewDB()

			// migrator, err := migrations.Init(db)
			// if err != nil {
			// 	fmt.Println("Unable to fetch migrator")
			// 	return
			// }

			// if err := migrator.MigrationStatus(); err != nil {
			// 	fmt.Println("Unable to fetch migration status")
			// 	return
			// }

			// return
		},
	}
}
