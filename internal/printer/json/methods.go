package json

import (
	"encoding/json"
	"fmt"

	"github.com/godbus/dbus/v5"
)

func (p *impl) Variant(v *dbus.Variant) {
	var err error
	var raw json.RawMessage
	raw, err = json.Marshal(v.Value())
	if err != nil {
		panic(err)
	}

	result := struct {
		Signature string          `json:"signature"`
		Value     json.RawMessage `json:"value"`
	}{
		Signature: v.Signature().String(),
		Value:     raw,
	}

	var resultJSON []byte
	if p.indent == "" {
		resultJSON, err = json.Marshal(result)
	} else {
		resultJSON, err = json.MarshalIndent(result, "", p.indent)
	}
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", resultJSON)
}
func (p *impl) Message(*dbus.Message) {
	panic("not implement yet")
}
