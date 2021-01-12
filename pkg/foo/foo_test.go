package foo_test

import (
	"github.com/frncmx/playground-travis/pkg/foo"
	"testing"
)

func TestBar(t *testing.T) {
	want := "bar"
	got := foo.Bar()

	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
