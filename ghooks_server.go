package ghooks

import (
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	Conf *Conf
}

func NewServer(conf *Conf) *Server {
	return &Server{conf}
}

var hooks Hooks

func (s *Server) Run() error {
	fmt.Printf("ghooks server start 127.0.0.1:%d \n", s.Conf.Port)
	http.HandleFunc("/", EventReciver)
	return http.ListenAndServe(":"+strconv.Itoa(s.Conf.Port), nil)
}

func EventReciver(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		http.Error(w, "Method Not Allowd", http.StatusMethodNotAllowed)
		return
	}

	event := req.Header.Get("X-GitHub-Event")

	if event == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	Emmit(event, req)
}
