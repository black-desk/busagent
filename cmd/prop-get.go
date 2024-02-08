package cmd

import (
	"github.com/godbus/dbus/v5"
	"github.com/spf13/cobra"
)

// propGetCmd represents the prop get command
var propGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get property of DBus object",
	RunE: func(cmd *cobra.Command, args []string) error {
		return agent.PropGet(
			flagName,
			dbus.ObjectPath(flagObjectPath),
			flagPropName)
	},
}

func init() {
	propCmd.AddCommand(propGetCmd)
}
