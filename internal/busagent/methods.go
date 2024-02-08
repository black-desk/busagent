package busagent

import (
	"errors"
	"strings"

	. "github.com/black-desk/lib/go/errwrap"
	"github.com/godbus/dbus/v5"
)

func (agent *impl) PropGet(
	id string, path dbus.ObjectPath, interfaceName string, propName string,
) (
	err error,
) {
	defer Wrap(&err, "get prop")

	obj := agent.conn.Object(id, path)

	var prop dbus.Variant
	prop, err = obj.GetProperty(interfaceName + "." + propName)
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

func (agent *impl) Call(
	id string, path dbus.ObjectPath, interfaceName string, method string,
	rawArgs ...string,
) (
	err error,
) {
	var args []any
	for i := 0; i < len(rawArgs); i++ {
		var v dbus.Variant
		v, err = dbus.ParseVariant(rawArgs[i], dbus.Signature{})
		if err != nil {
			return err
		}
		args = append(args, v.Value())
	}

	obj := agent.conn.Object(id, path)
	c := make(chan *dbus.Call, 1)
	call := obj.Go(interfaceName+"."+method, 0, c, args...)
	<-c
	if call.Err != nil {
		return call.Err
	}
	agent.print.Reply(call.Body)
	return
}
