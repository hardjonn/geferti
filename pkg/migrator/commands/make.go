package commands

import (
	"fmt"

	"github.com/hardjonn/geferti/migrations"

	"github.com/spf13/cobra"
)

func makeCmd(s migrations.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "make",
		Short: "make a new empty migrations file with the given name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]

			fmt.Println(name)

			s.Make(name)

			// name, err := cmd.Flags().GetString("name")
			// if err != nil {
			// 	fmt.Println("Unable to read flag `name`", err.Error())
			// 	return
			// }

			// if err := migrations.Create(name); err != nil {
			// 	fmt.Println("Unable to create migration", err.Error())
			// 	return
			// }
		},
	}
}
