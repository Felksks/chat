package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type hi struct {
	msg []string
}
type message struct {
	nickname string
	text     string
}

func (h *hi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.msg = append(h.msg, r.RequestURI)
	io.WriteString(w, strings.Join(h.msg, "\n"))

}
func (m *message) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
func main() {
	h := new(hi)
	h.msg = append(h.msg, "Hi, Feliksks!")
	http.Handle("/", h)
	http.Handle("/sendmessage", new(message))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
