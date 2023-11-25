package window

import (
	"github.com/fiatjaf/quickjs-go"
)

func InjectTo(ctx *quickjs.Context) error {
	ret, err := ctx.Eval(windowJs)
	defer ret.Free()

	if err != nil {
		return err
	}

	return nil
}
