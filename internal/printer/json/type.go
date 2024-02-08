package json

import (
	"github.com/black-desk/busagent/internal/printer"
	"go.uber.org/zap"
)

type impl struct {
	log    *zap.SugaredLogger
	prefix string
	indent string
}

var _ printer.Printer = &impl{}
