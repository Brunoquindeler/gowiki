package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	p := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	_ = httptest.NewRequest(http.MethodGet, "/view/TestPage", nil)
	w := httptest.NewRecorder()
	renderTemplate(w, "view", p)
}
