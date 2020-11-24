package service

import (
	"github.com/valyala/fasthttp"
	"fmt"
)


func Http_run(){
	m := func(ctx *fasthttp.RequestCtx) {
		headerString := ""
		ctx.Request.Header.VisitAll(func (key, value []byte) {
    		headerString += fmt.Sprintf("%s: %s\n",string(key), string(value))
		})
		Save("http",ctx.RemoteAddr().String(),headerString)
		ctx.Error("not found", fasthttp.StatusNotFound)
	}
	fasthttp.ListenAndServe(":5001", m)
}