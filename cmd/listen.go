package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen to dbus signal",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(flagMatchOption) == 0 {
			err = errors.New("match option is required")
			return
		}
		err = agent.Listen(flagMatchOption)
		return
	},
}

var flagMatchOption []string

func init() {
	rootCmd.AddCommand(listenCmd)
	listenCmd.PersistentFlags().StringArrayVarP(
		&flagMatchOption, "match", "m", []string{}, "match expression")
}
