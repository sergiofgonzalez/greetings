package greetings_test

import (
	"testing"

	"github.com/sergiofgonzalez/greetings"
)

func TestHello(t *testing.T) {
	want := "Hello, World!"
	got := greetings.Hello()

	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}


func TestHelloToJason(t *testing.T) {
	want := "Hello to Jason Isaacs!"
	got := greetings.HelloToJason()

	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}


func TestHelloTo(t *testing.T) {
	type testCase struct {
		name							string
		expectedGreeting	string
	}

	tests := map[string]testCase{
		"Name given": {
			name: "Idris",
			expectedGreeting: "Hello to Idris!",
		},
		"Empty name": {
			name: "",
			expectedGreeting: "Hello to stranger!"},
	}

	for subTestLabel, subTestCase := range tests {
		t.Run(subTestLabel, func(t *testing.T) {
			want := subTestCase.expectedGreeting

			got := greetings.HelloTo(subTestCase.name)
			if got != want {
				t.Errorf("%v: got %q, but wanted %q", subTestLabel, got, want)
			}
		})
	}
}