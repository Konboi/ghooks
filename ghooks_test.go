package ghooks

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var count int

func Push(paylod interface{}) {
	count++
}

func PullRequest(paylod interface{}) {
	count += 2
}

func TestEmmit(t *testing.T) {
	On("push", Push)
	On("pull_request", PullRequest)

	var payload interface{}
	Emmit("push", payload)

	if count != 1 {
		t.Fatal("Not call push Event")
	}

	Emmit("pull_request", payload)
	if count != 3 {
		t.Fatal("Not call pull_request Event")
	}

}

func TestReciver(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	Reciver(w, req)
	if w.Code == 200 {
		t.Fatalf("Allowd only POST Method but expected status 200; received %d", w.Code)
	}

	req, _ = http.NewRequest("POST", "/", nil)
	req.Header.Add("X-GitHub-Event", "")
	w = httptest.NewRecorder()
	Reciver(w, req)
	if w.Code == 200 {
		t.Fatalf("Event name is nil but return 200; received %d", w.Code)
	}

	req, _ = http.NewRequest("POST", "/", nil)
	req.Header.Set("X-GitHub-Event", "hoge")
	w = httptest.NewRecorder()
	Reciver(w, req)
	if w.Code == 200 {
		t.Fatalf("Body is nil but return 200; received %d", w.Code)
	}

	json_string := `{"fuga": "hoge", "foo": { "bar": "boo" }}`
	req, _ = http.NewRequest("POST", "/", strings.NewReader(json_string))
	req.Header.Set("X-GitHub-Event", "hoge")
	w = httptest.NewRecorder()
	Reciver(w, req)
	if w.Code != 200 {
		t.Fatalf("Not return 200; received %d", w.Code)
	}
}
