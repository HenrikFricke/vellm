package cmd

import (
	"errors"
	"fmt"
	"os"

	medium "github.com/Medium/medium-sdk-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var mediumClient *medium.Medium

var homeDir = os.Getenv("HOME")

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "vellm",
	Short: "CLI for Medium",
	Long:  `Write stories in Markdown and upload them in seconds to Medium.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".vellm") // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		mediumClient = medium.NewClientWithAccessToken(viper.GetString("MEDIUM_TOKEN"))
	}

}

func mediumTokenCheck(cmd *cobra.Command, args []string) error {
	if token = viper.GetString("MEDIUM_TOKEN"); token == "" {
		return errors.New("Please setup Vellm first to work with Medium. Run vellm setup.")
	}

	if _, err := mediumClient.GetUser(""); err != nil {
		return errors.New(err.Error() + ">> Run `vellm setup` if invalid token given.")
	}

	return nil
}
