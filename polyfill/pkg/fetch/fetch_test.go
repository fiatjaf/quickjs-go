package fetch_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fiatjaf/quickjs-go"
	"github.com/fiatjaf/quickjs-go/polyfill/pkg/fetch"
	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {

	rt := quickjs.NewRuntime()
	defer rt.Close()

	ctx := rt.NewContext()
	defer ctx.Close()

	fetch.InjectTo(ctx)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; utf-8")
		_, _ = w.Write([]byte(`{"status": true}`))
	}))

	val, _ := ctx.Eval(fmt.Sprintf(`
	var ret;
	fetch('%s', {Method: 'GET'}).then(response => response.json()).then(data => {
		ret = data.status;
		return data;
	})`, srv.URL))
	defer val.Free()

	for {
		_, err := rt.ExecutePendingJob()
		if err == io.EOF {
			err = nil
			break
		}
	}

	asyncRet, _ := ctx.Eval("ret")
	defer asyncRet.Free()

	require.EqualValues(t, true, asyncRet.Bool())

}
