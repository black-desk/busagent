package printer

import (
	. "github.com/black-desk/lib/go/errwrap"
	"go.uber.org/zap"
)

type Opt = (func(*impl) (*impl, error))

func New(opts ...Opt) (ret Printer, err error) {
	defer Wrap(&err, "create new printer")

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

	ret = c
	return
}

func WithLogger(logger *zap.SugaredLogger) Opt {
	if logger == nil {
		panic("nil logger")
	}

	return func(p *impl) (ret *impl, err error) {
		p.log = logger
		ret = p
		return
	}
}
