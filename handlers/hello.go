package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger // kinf of dependency injection
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello")

	data, err := ioutil.ReadAll(r.Body)
	h.l.Println(string(data))

	if err != nil {
		http.Error(rw, "Error occurred", http.StatusBadRequest)
		// rw.WriteHeader(http.StatusBadRequest) // another wat to do the same
		// rw.Write([]byte("Invalid"))
		return
	}

	rw.Write([]byte("Hello"))
}
