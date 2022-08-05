package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("test")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Ngaco", http.StatusBadRequest)
		return
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("ngaco"))
	}

	fmt.Fprintf(w, "hello %s\n", d)
}
