package json

import (
	"fmt"

	json "github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"

	"github.com/godbus/dbus/v5"
)

func (p *impl) Variant(variant *dbus.Variant) {
	v := struct {
		Signature dbus.Signature
		Value     interface{}
	}{
		Signature: variant.Signature(),
		Value:     variant.Value(),
	}
	var err error
	var raw []byte
	raw, err = json.Marshal(v, jsontext.WithIndent(p.indent))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", raw)
}

func (p *impl) Signal(s *dbus.Signal) {
	var err error
	var raw []byte
	raw, err = json.Marshal(s,
		jsontext.WithIndent(p.indent),
		json.WithMarshalers(json.MarshalFuncV1[*dbus.Variant](
			func(v *dbus.Variant) ([]byte, error) {
				return json.Marshal(struct {
					Signature string
					Value     interface{}
				}{
					Signature: v.Signature().String(),
					Value:     v.Value()})
			}),
		))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", raw)
}

func (p *impl) Reply(c *dbus.Call) {
	v := struct {
		Destination      string
		Path             dbus.ObjectPath
		Method           string
		Args             []interface{}
		Err              error
		Body             []interface{}
		ResponseSequence dbus.Sequence
	}{
		Destination:      c.Destination,
		Path:             c.Path,
		Method:           c.Method,
		Args:             c.Args,
		Err:              c.Err,
		Body:             c.Body,
		ResponseSequence: c.ResponseSequence,
	}
	var err error
	var raw []byte
	raw, err = json.Marshal(v, jsontext.WithIndent(p.indent))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", raw)
}
