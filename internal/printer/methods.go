package printer

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

func (p *impl) Variant(v *dbus.Variant) {
	fmt.Printf("%s %s\n", v.Signature(), v.String())
}

func (p *impl) Signal(s *dbus.Signal) {
	fmt.Printf("%v\n", *s)
}
