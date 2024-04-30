package main

import "testing"

func TestHello(t *testing.T) {
	greetString := "harshit"
	got := Hello(greetString)
	want := "Hello, " + greetString

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
