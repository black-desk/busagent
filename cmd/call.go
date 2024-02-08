/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/spf13/cobra"
)

// callCmd represents the call command
var callCmd = &cobra.Command{
	Use:   "call",
	Short: "Call dbus method",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if flagName == "" {
			err = fmt.Errorf(`"name" is required`)
			return
		}

		if flagInterface == "" {
			err = fmt.Errorf(`"interface" is required`)
			return
		}

		if flagMethodName == "" {
			err = fmt.Errorf(`"method" is required`)
			return
		}

		err = agent.Call(
			flagName,
			dbus.ObjectPath(flagObjectPath),
			flagInterface,
			flagMethodName,
			args...)

		return err
	},
}

var flagMethodName string

func init() {
	rootCmd.AddCommand(callCmd)
	callCmd.PersistentFlags().StringVarP(
		&flagMethodName, "method", "m", "", "method name")
}
