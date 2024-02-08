package busagent

import (
	"github.com/black-desk/busagent/internal/printer"
	"github.com/godbus/dbus/v5"
	"go.uber.org/zap"
)

//go:generate go run github.com/rjeczalik/interfaces/cmd/interfacer@v0.3.0 -for github.com/black-desk/busagent/internal/busagent.impl -as busagent.BusAgent -o interface.go
type impl struct {
	print printer.Printer
	log   *zap.SugaredLogger
	conn  *dbus.Conn
}
