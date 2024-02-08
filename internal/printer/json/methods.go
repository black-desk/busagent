package json

import (
	"encoding/json"
	"fmt"

	"github.com/godbus/dbus/v5"
)

func (p *impl) Variant(variant *dbus.Variant) {
	v := struct {
		Signature dbus.Signature
		Value     interface{}
	}{
		Signature: variant.Signature(),
		Value: variant.Value(),
	}
	var err error
	var raw json.RawMessage
	if p.indent == "" {
		raw, err = json.Marshal(v)
	} else {
		raw, err = json.MarshalIndent(v, "", p.indent)
	}
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", raw)
}

func (p *impl) Signal(s *dbus.Signal) {
	var err error
	var raw json.RawMessage
	if p.indent == "" {
		raw, err = json.Marshal(s)
	} else {
		raw, err = json.MarshalIndent(s, "", p.indent)
	}
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
	var raw json.RawMessage
	if p.indent == "" {
		raw, err = json.Marshal(v)
	} else {
		raw, err = json.MarshalIndent(v, "", p.indent)
	}
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", raw)
}
