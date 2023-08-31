package handler

import "net/http"


type ping struct{}

func NewPing() *ping {
	return &ping{}
}

func (h *ping) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG!"))
	w.WriteHeader(http.StatusOK)
}
