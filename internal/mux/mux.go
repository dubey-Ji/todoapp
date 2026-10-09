package mux

import (
	"net/http"

	"github.com/dubey-Ji/todoapp/internal/hanlders"
)

func Mux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", hanlders.Healthz)

	return mux
}
