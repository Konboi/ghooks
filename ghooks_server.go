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

func (s *Server) Run() error {
	fmt.Printf("ghooks server start port:%d \n", s.Conf.Port)
	http.HandleFunc("/", EventReciver)
	return http.ListenAndServe(":"+strconv.Itoa(s.Conf.Port), nil)
}

func EventReciver(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "fooooooooooo")
}
