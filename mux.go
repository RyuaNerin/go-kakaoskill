package skill

import "net/http"

// ServeMux 핸들러 처리를 편하게 해주는 유틸리티
type MuxHelper struct {
	remoteAddrHeader string

	mux *http.ServeMux
}

// mux : nil 이면 http.DefaultServeMux 를 사용합니다.
func NewMuxHelper(mux *http.ServeMux, remoteAddrHeaderName string) *MuxHelper {
	if mux == nil {
		mux = http.DefaultServeMux
	}

	return &MuxHelper{
		remoteAddrHeader: remoteAddrHeaderName,
		mux:              mux,
	}
}

func (m *MuxHelper) F(pattern string, fn HandleFunc) {
	if len(m.remoteAddrHeader) == 0 {
		m.mux.HandleFunc(pattern, F(fn))
	} else {
		m.mux.HandleFunc(pattern, FP(m.remoteAddrHeader, fn))
	}
}

func (m *MuxHelper) H(pattern string, h Handler) {
	if len(m.remoteAddrHeader) == 0 {
		m.mux.HandleFunc(pattern, H(h))
	} else {
		m.mux.HandleFunc(pattern, HP(m.remoteAddrHeader, h))
	}
}
