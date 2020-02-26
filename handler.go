package skill

import (
	"encoding/json"
	"net"
	"net/http"
)

// 카카오톡 스킬 서버의 IP 입니다.
var AllowedRemoteAddr = []string{
	"219.249.231.40",
	"219.249.231.41",
	"219.249.231.42",
}

type Handler interface {
	Handle(ctx *Context)
}

type HandleFunc func(ctx *Context)

func (hf HandleFunc) Handle(ctx *Context) {
	hf(ctx)
}

func FP(headerName string, fn HandleFunc) http.HandlerFunc {
	return handle(headerName, fn)
}

func F(fn HandleFunc) http.HandlerFunc {
	return handle("", fn)
}

func HP(headerName string, h Handler) http.HandlerFunc {
	return handle(headerName, h)
}

func H(h Handler) http.HandlerFunc {
	return handle("", h)
}

func handle(headerName string, handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		allowed := false

		reqAddr := r.RemoteAddr
		if headerName != "" {
			if headerValue := r.Header.Get(headerName); headerValue != "" {
				reqAddr = headerValue
			}
		}

		if host, _, err := net.SplitHostPort(reqAddr); err == nil {
			reqAddr = host
		}

		if len(AllowedRemoteAddr) == 0 {
			allowed = true
		} else {
			for _, addr := range AllowedRemoteAddr {
				if addr == reqAddr {
					allowed = true
					break
				}
			}
		}

		if !allowed {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}

		var payload SkillPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			panic(err)
		}

		sctx := Context{
			w:       w,
			Reqeust: r,
			Payload: &payload,
		}
		handler.Handle(&sctx)

		if !sctx.sent {
			w.WriteHeader(http.StatusNoContent)
			panic(ErrorNoResponse)
		}
	}
}