package commands

import (
	"github.com/hardjonn/geferti/migrations"

	"github.com/spf13/cobra"
)

func downCmd(s migrations.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "run down migrations",
		Run: func(cmd *cobra.Command, args []string) {

			// step, err := cmd.Flags().GetInt("step")
			// if err != nil {
			// 	fmt.Println("Unable to read flag `step`")
			// 	return
			// }

			// db := app.NewDB()

			// migrator, err := migrations.Init(db)
			// if err != nil {
			// 	fmt.Println("Unable to fetch migrator")
			// 	return
			// }

			// err = migrator.Down(step)
			// if err != nil {
			// 	fmt.Println("Unable to run `down` migrations")
			// 	return
			// }
		},
	}
}
