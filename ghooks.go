package ghooks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	VERSION = 0.1
)

type Server struct {
	Port int
}

type Hook struct {
	Event string
	Func  func(payload interface{})
}

type Hooks struct {
	Hooks []Hook
}

var hooks Hooks

func (s *Server) On(name string, handler func(payload interface{})) {
	hooks.Hooks = append(hooks.Hooks, Hook{Event: name, Func: handler})
}

func Emmit(name string, payload interface{}) {
	for _, v := range hooks.Hooks {
		if strings.EqualFold(v.Event, name) {
			v.Func(payload)
		}
	}
}

func NewServer(port int) *Server {
	return &Server{port}
}

func (s *Server) Run() error {
	fmt.Printf("ghooks server start 0.0.0.0:%d \n", s.Port)
	http.HandleFunc("/", Reciver)
	return http.ListenAndServe(":"+strconv.Itoa(s.Port), nil)
}

func Reciver(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		http.Error(w, "Method Not Allowd", http.StatusMethodNotAllowed)
		return
	}

	event := req.Header.Get("X-GitHub-Event")

	if event == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if req.Body == nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var payload interface{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	Emmit(event, payload)
	w.WriteHeader(http.StatusOK)
}
