# Polyfill for [quickjs-go](https://github.com/fiatjaf/quickjs-go)
[English](README.md) | 简体中文

## 功能
* fetch: `fetch`
* base64: `atob` and `btoa`
* window: `window`
* console: `console.log` and `console.error` and `console.warn` and `console.info` and `console.debug` and `console.trace`
* timers: `setTimeout` and `setInterval` and `clearTimeout` and `clearInterval`

### 用法
```go
package main
import (
	"time"

	"github.com/fiatjaf/quickjs-go"
	"github.com/fiatjaf/quickjs-go/polyfill"
)

func main() {
	// Create a new runtime
	rt := quickjs.NewRuntime()
	defer rt.Close()

	// Create a new context
	ctx := rt.NewContext()
	defer ctx.Close()

	// Inject polyfills to the context
	polyfill.InjectAll(ctx)

	ret, _ := ctx.Eval(`
	setTimeout(() => {
		fetch('https://api.github.com/users/buke', {Method: 'GET'}).then(response => response.json()).then(data => {
			console.log(data.login);
		});
	}, 50);
	`)
	defer ret.Free()

	// Wait for the timeout to finish
	time.Sleep(time.Millisecond * 100)

	rt.ExecuteAllPendingJobs()

	// Output:
	// buke
}
```

## 版权协议
[MIT](./LICENSE)
