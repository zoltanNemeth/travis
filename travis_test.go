package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 2
	actual := Add(1, 1)
	if actual != expected {
		t.Errorf("expected 2 and got %v", actual)
	}
}

func TestRemove(t *testing.T) {
	expected := 2
	actual := Remove(3, 1)
	if actual != expected {
		t.Errorf("expected 2 and got %v", actual)
	}
}
