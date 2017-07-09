package cmd

import (
	"fmt"

	medium "github.com/Medium/medium-sdk-go"
	"github.com/spf13/cobra"
)

// meCmd represents the me command
var meCmd = &cobra.Command{
	Use:     "me",
	Short:   "Get information about your Medium profile.",
	Long:    `Get information about your Medium profile.`,
	PreRunE: mediumTokenCheck,
	RunE: func(cmd *cobra.Command, args []string) error {
		var u *medium.User
		var err error

		if u, err = mediumClient.GetUser(""); err != nil {
			return err
		}

		fmt.Printf("I've found something about you: \n\n")
		fmt.Println("Username: " + u.Username)
		fmt.Println("Name: " + u.Name)
		fmt.Println("URL: " + u.URL)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(meCmd)
}
