package busagent

import (
	"errors"
	"strings"

	. "github.com/black-desk/lib/go/errwrap"
	"github.com/godbus/dbus/v5"
)

func (agent *impl) PropGet(
	id string, path dbus.ObjectPath, propName string,
) (
	err error,
) {
	defer Wrap(&err, "get prop")

	obj := agent.conn.Object(id, path)

	var prop dbus.Variant
	prop, err = obj.GetProperty(propName)
	if err != nil {
		return
	}

	agent.print.Variant(&prop)
	return
}

func (agent *impl) Listen(options []string) (err error) {
	defer Wrap(&err, "listen")
	opts := []dbus.MatchOption{}
	for i := range options {
		agent.log.Debug(options[i])
		pos := strings.Index(options[i], "=")
		if pos == -1 {
			return errors.New("invalid option " + options[i])
		}
		key := options[i][:pos]
		value := options[i][pos+1:]
		opts = append(opts, dbus.WithMatchOption(key, value))
	}

	err = agent.conn.AddMatchSignal(opts...)
	if err != nil {
		return
	}

	signals := make(chan *dbus.Signal)

	agent.conn.Signal(signals)

	for signal := range signals {
		agent.print.Signal(signal)
	}

	return
}
