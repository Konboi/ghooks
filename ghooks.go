package ghooks

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	VERSION = 0.2
)

type Server struct {
	Port   int
	Secret string
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
	return &Server{Port: port}
}

func (s *Server) Run() error {
	fmt.Printf("ghooks server start 0.0.0.0:%d \n", s.Port)
	http.HandleFunc("/", s.Reciver)
	return http.ListenAndServe(":"+strconv.Itoa(s.Port), nil)
}

func (s *Server) Reciver(w http.ResponseWriter, req *http.Request) {
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

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if s.Secret != "" {
		signature := req.Header.Get("X-Hub-Signature")
		if !s.isValidSignature(body, signature) {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	}

	var payload interface{}
	var decoder *json.Decoder

	if strings.Contains(req.Header.Get("Content-Type"), "application/json") {

		decoder = json.NewDecoder(bytes.NewReader(body))

	} else if strings.Contains(req.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {

		v, err := url.ParseQuery(string(body))
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		p := v.Get("payload")
		decoder = json.NewDecoder(strings.NewReader(p))
	}

	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	Emmit(event, payload)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) isValidSignature(body []byte, signature string) bool {

	if !strings.HasPrefix(signature, "sha1=") {
		return false
	}

	mac := hmac.New(sha1.New, []byte(s.Secret))
	mac.Write(body)
	actual := mac.Sum(nil)

	expected, err := hex.DecodeString(signature[5:])
	if err != nil {
		return false
	}

	return hmac.Equal(actual, expected)
}
