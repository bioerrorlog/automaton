package cmd

import (
	"testing"
)

func TestInit1d(t *testing.T) {
	got := len(Init1d(100))
	want := 100
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
