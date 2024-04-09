package set

import (
	"testing"
)

func equals(a, b Set[string]) bool {
	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}
	for k := range b {
		if _, ok := a[k]; !ok {
			return false
		}
	}
	return true
}

func TestAdd(t *testing.T) {

	got := Make[string]()
	got.Add("foo")
	got.Add("bar")

	want := map[string]struct{}{
		"foo": {},
		"bar": {},
	}

	if !equals(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestContains(t *testing.T) {
	got := Make[string]()
	got.Add("foo")

	if !got.Contains("foo") {
		t.Errorf("expect %q to contain %v", got, "foo")
	}
	if got.Contains("bar") {
		t.Errorf("expect %q not to contain %v", got, "bar")
	}
}

func TestDelete(t *testing.T) {
	got := Make[string]()
	got.Add("foo")
	got.Add("bar")
	got.Delete("bar")
	got.Delete("not present")

	want := map[string]struct{}{
		"foo": {},
	}

	if !equals(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
