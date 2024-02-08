package cmd

import (
	"fmt"
	"os"

	"github.com/black-desk/busagent/internal/busagent"
	"github.com/black-desk/busagent/internal/printer"
	"github.com/black-desk/busagent/internal/printer/json"
	"github.com/black-desk/lib/go/logger"
	"github.com/godbus/dbus/v5"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var agent busagent.BusAgent

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "busagent",
	Short: "Yet another DBus utility CLI written in golang.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		log := zap.NewNop().Sugar()
		if flagVerbose {
			log = logger.Get("busagent")
		}

		var resultPrinter printer.Printer
		if flagJSON {
			resultPrinter, err = json.New(
				json.WithLogger(log),
				json.WithIndent(flagIndent),
			)
		} else {
			resultPrinter, err = printer.New(
				printer.WithLogger(log),
			)
		}
		if err != nil {
			return
		}

		var conn *dbus.Conn
		if flagBusAddress != "" {
			conn, err = dbus.Dial(flagBusAddress)
		} else if flagBusType == "session" {
			conn, err = dbus.ConnectSessionBus()
		} else if flagBusType == "system" {
			conn, err = dbus.ConnectSystemBus()
		} else {
			err = fmt.Errorf("unknow bus type %s", flagBusType)
		}
		if err != nil {
			return
		}

		agent, err = busagent.New(
			busagent.WithLogger(log),
			busagent.WithPrinter(resultPrinter),
			busagent.WithDBusConn(conn),
		)
		if err != nil {
			return
		}

		return
	},
}

func Execute() {
	cobra.EnableTraverseRunHooks = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var flagVerbose bool
var flagJSON bool
var flagBusType string
var flagBusAddress string
var flagIndent string

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&flagVerbose, "verbose", "v", false, "enable log")
	rootCmd.PersistentFlags().BoolVarP(
		&flagJSON, "json", "j", false, "output in json")
	rootCmd.PersistentFlags().StringVarP(
		&flagBusType, "type", "t", "session", `well know message bus to use, "session" or "system"`)
	rootCmd.PersistentFlags().StringVarP(
		&flagBusAddress, "addr", "a", "", "message bus address to use, this ignore --type")
	rootCmd.PersistentFlags().StringVarP(
		&flagIndent, "indent", "", "  ", "json indent")
}
