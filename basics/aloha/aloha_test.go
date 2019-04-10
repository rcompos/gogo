package main

import "testing"

func TestWhom(t *testing.T) {
	whom := aloha()
	if whom != "World" {
		t.Error("Unexpected response from aloha")
	}
}
