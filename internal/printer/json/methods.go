package json

import (
	"fmt"

	json "github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"

	"github.com/godbus/dbus/v5"
)

func marshalDBusSignature(s dbus.Signature) ([]byte, error) {
	return json.Marshal(s.String())
}

func marshalDBusVariant(v *dbus.Variant) ([]byte, error) {
	return json.Marshal(
		struct {
			Signature dbus.Signature
			Value     interface{}
		}{
			Signature: v.Signature(),
			Value:     v.Value(),
		},
		json.WithMarshalers(json.MarshalFuncV1(marshalDBusSignature)),
	)
}

var marshallers = json.NewMarshalers(
	json.MarshalFuncV1(marshalDBusSignature),
	json.MarshalFuncV1(marshalDBusVariant),
)

func (p *impl) Variant(v *dbus.Variant) {
	var err error
	var raw []byte
	raw, err = json.Marshal(
		v,
		jsontext.WithIndent(p.indent),
		json.WithMarshalers(marshallers),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", raw)
}

func (p *impl) Signal(s *dbus.Signal) {
	var err error
	var raw []byte
	raw, err = json.Marshal(
		s,
		jsontext.WithIndent(p.indent),
		json.WithMarshalers(json.MarshalFuncV1(marshalDBusVariant)))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", raw)
}

func (p *impl) Reply(c *dbus.Call) {
	var err error
	var raw []byte
	raw, err = json.Marshal(
		c,
		jsontext.WithIndent(p.indent),
		json.WithMarshalers(json.MarshalFuncV1(marshalDBusVariant)),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", raw)
}
