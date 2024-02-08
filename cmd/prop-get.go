package cmd

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/spf13/cobra"
)

// propGetCmd represents the prop get command
var propGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get property of DBus object",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if flagName == "" {
			err = fmt.Errorf(`"name" is required`)
			return
		}

		if flagInterface == "" {
			err = fmt.Errorf(`"prop" is required`)
			return
		}
		err = agent.PropGet(
			flagName,
			dbus.ObjectPath(flagObjectPath),
			flagInterface,
			flagPropName)
		return
	},
}

func init() {
	propCmd.AddCommand(propGetCmd)
}
