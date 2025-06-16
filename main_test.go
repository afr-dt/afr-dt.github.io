package main

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	want := "<h1>Hi 👋, I'm Alejandro</h1>"
	if got := Render(); !strings.Contains(got, want) {
		t.Fatalf("Render() no contiene %q", want)
	}
}
