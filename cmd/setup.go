package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var setupDescription = `Vellm needs an integration token of Medium to publish stories. 
Go to Medium and create a new token, go back to vellm and add the token.

Medium settings: https://medium.com/me/settings
`

var token string

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup Vellm to use it with Medium",
	Long:  setupDescription,
	RunE: func(cmd *cobra.Command, args []string) error {
		if token != "" {
			return storeToken(token)
		}

		fmt.Println(setupDescription)

		// read token from input
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter integration token: ")
		enteredToken, err := reader.ReadString('\n')

		if err != nil {
			return err
		}

		if len(enteredToken) < 10 {
			return errors.New("Not a valid token. Please enter a proper value.")
		}

		return storeToken(enteredToken)
	},
}

func storeToken(token string) error {
	d1 := []byte(fmt.Sprintf("MEDIUM_TOKEN: %s", token))
	err := ioutil.WriteFile(fmt.Sprintf("%s/.vellm.yaml", homeDir), d1, 0644)

	if err != nil {
		return errors.New("Couldn't store token...")
	}

	fmt.Println("Token was stored. You can now start using Vellm :)")
	return nil
}

func init() {
	RootCmd.AddCommand(setupCmd)

	setupCmd.Flags().StringVarP(&token, "token", "t", "", "Medium integration token to grant access for Vellm")
}
