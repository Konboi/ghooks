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
	req.Header.Add("X-GitHub-Event", "hoge")
	w := httptest.NewRecorder()
	EventReciver(w, req)
	if w.Code != 200 {
		t.Fatalf("expected status 200; received %d", w.Code)
	}
}
