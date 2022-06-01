package main

import (
	"bytes"
	"testing"
)

func TestSaveAndLoadPage(t *testing.T) {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}

	err := p1.save()
	if err != nil {
		t.Fatal(err)
	}

	p2, _ := loadPage("TestPage")
	if !bytes.Equal(p1.Body, p2.Body) {
		t.Errorf("%v and %v should be equal", p1.Body, p2.Body)
	}
}

func TestLoadPageErr(t *testing.T) {
	_, err := loadPage("Test Page")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
