// Code generated by interfacer; DO NOT EDIT

package printer

import (
	"github.com/godbus/dbus/v5"
)

// Printer is an interface generated for "github.com/black-desk/busagent/internal/printer.impl".
type Printer interface {
	Reply(*dbus.Call)
	Signal(*dbus.Signal)
	Variant(*dbus.Variant)
}
