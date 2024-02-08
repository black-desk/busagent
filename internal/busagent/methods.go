package busagent

import (
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
