// example.go
package main

import (
	"net/http"

	skill "github.com/RyuaNerin/go-kakaoskill/v2"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/simple1", skill.F(simpleFunc))
	mux.Handle("/simple2", skill.H(new(simpleHandler)))

	server := http.Server{
		Handler: mux,
	}

	server.ListenAndServe()
}

func simpleFunc(ctx *skill.Context) {
	ctx.WriteSimpleText("simple text")
}

type simpleHandler int

func (h simpleHandler) Handle(ctx *skill.Context) {
	ctx.WriteSimpleText("simple text")
}
