package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// propCmd represents the prop command
var propCmd = &cobra.Command{
	Use:   "prop",
	Short: "Get or set property of DBus object",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if flagName == "" {
			err = fmt.Errorf(`"name" is required`)
			return
		}

		if flagPropName == "" {
			err = fmt.Errorf(`"prop" is required`)
			return
		}

		return
	},
}

var flagName string
var flagObjectPath string
var flagPropName string

func init() {
	rootCmd.AddCommand(propCmd)
	propCmd.PersistentFlags().StringVarP(
		&flagName, "name", "n", "", "DBus name")
	propCmd.PersistentFlags().StringVarP(
		&flagObjectPath, "path", "o", "/", "DBus object path")
	propCmd.PersistentFlags().StringVarP(
		&flagPropName, "prop", "p", "", "property name")
}
