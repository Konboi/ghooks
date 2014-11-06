package ghooks

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEventReciver(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("X-GitHub-Event", "push")
	w := httptest.NewRecorder()
	EventReciver(w, req)
	if w.Code != 200 {
		t.Fatalf("expected status 200; received %d", w.Code)
	}

	req.Header.Set("X-GitHub-Event", "")
	w = httptest.NewRecorder()
	EventReciver(w, req)
	if w.Code == 200 {
		t.Fatalf("Event name is nil but return 200; received %d", w.Code)
	}

	req, _ = http.NewRequest("GET", "/", nil)
	req.Header.Add("X-GitHub-Event", "push")
	w = httptest.NewRecorder()
	EventReciver(w, req)
	if w.Code == 200 {
		t.Fatalf("Allowd only POST Method but expected status 200; received %d", w.Code)
	}
}
