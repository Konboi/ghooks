package ghooks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var count int
var some_string string

func Push(paylod interface{}) {
	count++
}

func Push2(payload interface{}) {
	some_string = payload.(map[string]interface{})["fuga"].(string)
}

func PullRequest(paylod interface{}) {
	count += 2
}

func TestNewServer(t *testing.T) {
	hooks := NewServer(999999)

	if hooks.Host != "0.0.0.0" {
		t.Fatal("Default host is 0.0.0.0")
	}

	hooks2 := NewServer(999999, "127.0.0.1")

	if hooks2.Host != "127.0.0.1" {
		t.Fatalf("Not equal 127.0.0.1")
	}
}

func TestEmit(t *testing.T) {
	hooks := NewServer(999999)
	hooks.On("push", Push)
	hooks.On("pull_request", PullRequest)
	hooks.On("push2", Push2)

	var payload interface{}
	Emit("push", payload)

	if count != 1 {
		t.Fatal("Not call push Event")
	}

	Emit("pull_request", payload)
	if count != 3 {
		t.Fatal("Not call pull_request Event")

	}

	b := []byte(`{"fuga": "hoge"}`)
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.Decode(&payload)
	Emit("push2", payload)
	if !strings.EqualFold(some_string, "hoge") {
		t.Fatal("Cannot  access payload")
	}

}

func TestReceiver(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	s := &Server{}
	s.Receiver(w, req)
	if w.Code == 200 {
		t.Fatalf("Allowd only POST Method but expected status 200; received %d", w.Code)
	}

	req, _ = http.NewRequest("POST", "/", nil)
	req.Header.Add("X-GitHub-Event", "")
	w = httptest.NewRecorder()
	s = &Server{}
	s.Receiver(w, req)
	if w.Code == 200 {
		t.Fatalf("Event name is nil but return 200; received %d", w.Code)
	}

	req, _ = http.NewRequest("POST", "/", nil)
	req.Header.Set("X-GitHub-Event", "hoge")
	w = httptest.NewRecorder()
	s = &Server{}
	s.Receiver(w, req)
	if w.Code == 200 {
		t.Fatalf("Body is nil but return 200; received %d", w.Code)
	}

	json_string := `{"fuga": "hoge", "foo": { "bar": "boo" }}`
	req, _ = http.NewRequest("POST", "/", strings.NewReader(json_string))
	req.Header.Set("X-GitHub-Event", "hoge")
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	s = &Server{}
	s.Receiver(w, req)
	if w.Code != 200 {
		t.Fatalf("Not return 200; received %d", w.Code)
	}

	json_string = `{"fuga": "hoge", "foo": { "bar", "boo" }}`
	req, _ = http.NewRequest("POST", "/", strings.NewReader(json_string))
	req.Header.Set("X-GitHub-Event", "hoge")
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	s = &Server{}
	s.Receiver(w, req)
	if w.Code == 200 {
		t.Fatalf("Should not be 200; received %d", w.Code)
	}

	json_string = `{"fuga": "hoge", "foo": { "bar": "boo" }}`
	req, _ = http.NewRequest("POST", "/", strings.NewReader("payload="+json_string))
	req.Header.Set("X-GitHub-Event", "hoge")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	s = &Server{}
	s.Receiver(w, req)
	if w.Code != 200 {
		t.Fatalf("Not return 200; received %d", w.Code)
	}

	json_string = `{"fuga": "hoge", "foo": { "bar": "boo" }}`
	req, _ = http.NewRequest("POST", "/", strings.NewReader(json_string))
	req.Header.Set("X-GitHub-Event", "hoge")
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	s = &Server{Secret: "mysecret"}
	s.Receiver(w, req)
	if w.Code != 400 {
		t.Fatalf("Not return 400; received %d", w.Code)
	}

	json_string = `{"fuga": "hoge", "foo": { "bar": "boo" }}`
	req, _ = http.NewRequest("POST", "/", strings.NewReader(json_string))
	req.Header.Set("X-GitHub-Event", "hoge")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hub-Signature", "sha1=invalid")
	w = httptest.NewRecorder()
	s = &Server{Secret: "dameleon"}
	s.Receiver(w, req)
	if w.Code != 400 {
		t.Fatalf("Not return 400; received %d", w.Code)
	}

	json_string = `{"fuga": "hoge", "foo": { "bar": "boo" }}`
	req, _ = http.NewRequest("POST", "/", strings.NewReader(json_string))
	req.Header.Set("X-GitHub-Event", "hoge")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hub-Signature", "sha1=17f693f6f260c0e4b4090ae1e0cf195e03bed614")
	w = httptest.NewRecorder()
	s = &Server{Secret: "mysecret"}
	s.Receiver(w, req)
	if w.Code != 200 {
		t.Fatalf("Not return 200; received %d", w.Code)
	}
}
