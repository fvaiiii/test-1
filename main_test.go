package main

import "testing"

type TestCase struct {
	name  string
	value string
}

func TestHello(t *testing.T) {
	cases := TestCase{
		name:  "1test",
		value: "Hello, CI/CD!",
	}
	t.Run(cases.name, func(t *testing.T) {
		res := hello()
		if res != cases.value {
			t.Errorf("value not equal")
		}
	})
}
