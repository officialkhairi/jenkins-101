package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Jenkins")
	want := "Hello, Jenkins!"
	if got != want {
		t.Errorf("Greet() = %q, want %q", got, want)
	}
}
