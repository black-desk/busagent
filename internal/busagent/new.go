package busagent

import (
	"github.com/black-desk/busagent/internal/printer"
	. "github.com/black-desk/lib/go/errwrap"
	"github.com/godbus/dbus/v5"
	"go.uber.org/zap"
)

type Opt = (func(*impl) (*impl, error))

func New(opts ...Opt) (ret BusAgent, err error) {
	defer Wrap(&err, "create new bus agent")

	c := &impl{}
	for i := range opts {
		c, err = opts[i](c)
		if err != nil {
			c = nil
			return
		}
	}

	if c.log == nil {
		panic("nil logger")
	}

	if c.print == nil {
		panic("nil print")
	}

	c.log.Debug("new bus agent created")

	ret = c
	return
}

func WithLogger(logger *zap.SugaredLogger) Opt {
	if logger == nil {
		panic("nil logger")
	}

	return func(agent *impl) (ret *impl, err error) {
		agent.log = logger
		ret = agent
		return
	}
}

func WithPrinter(p printer.Printer) Opt {
	if p == nil {
		panic("nil print")
	}

	return func(agent *impl) (ret *impl, err error) {
		agent.print = p
		ret = agent
		return
	}
}

func WithDBusConn(conn *dbus.Conn) Opt {
	if conn == nil {
		panic("nil conn")
	}

	return func(agent *impl) (ret *impl, err error) {
		agent.conn = conn
		ret = agent
		return
	}
}
