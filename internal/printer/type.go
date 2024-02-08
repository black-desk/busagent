package printer

import "go.uber.org/zap"

//go:generate go run github.com/rjeczalik/interfaces/cmd/interfacer@v0.3.0 -for github.com/black-desk/busagent/internal/printer.impl -as printer.Printer -o interface.go

type impl struct {
	log *zap.SugaredLogger
}

var _ Printer = &impl{}
