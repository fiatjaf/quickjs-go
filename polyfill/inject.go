package polyfill

import (
	"github.com/fiatjaf/quickjs-go"
	"github.com/fiatjaf/quickjs-go/polyfill/pkg/base64"
	"github.com/fiatjaf/quickjs-go/polyfill/pkg/console"
	"github.com/fiatjaf/quickjs-go/polyfill/pkg/fetch"
	"github.com/fiatjaf/quickjs-go/polyfill/pkg/timer"
	"github.com/fiatjaf/quickjs-go/polyfill/pkg/window"
)

func InjectAll(ctx *quickjs.Context) {
	window.InjectTo(ctx)
	fetch.InjectTo(ctx)
	console.InjectTo(ctx)
	base64.InjectTo(ctx)
	timer.InjectTo(ctx)
}
