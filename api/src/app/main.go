package app

import (
	"context"
	"net/http"

	"github.com/favclip/ucon"
)

func init() {
	//var _ ucon.HTTPErrorResponse = &HTTPError{}
	//ucon.Middleware(UseAppengineContext)
	//ucon.Middleware(DomainAuth)
	ucon.Orthodox()
	//ucon.Middleware(UseReqRespLogger)
	//ucon.Middleware(swagger.RequestValidator())

	ucon.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request, c context.Context) {
		w.Write([]byte("Hello World!"))
	})

	ucon.DefaultMux.Prepare()
	http.Handle("/", ucon.DefaultMux)
}
