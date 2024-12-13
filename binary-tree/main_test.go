package main

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestWalk(t *testing.T) {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := 1; i <= 10; i++ {
		got := <-ch
		if i != got {
			t.Errorf("expected: %d, actual: %d", i, got)
		}
	}
}

func TestSame(t *testing.T) {
	if !Same(tree.New(1), tree.New(1)) {
		t.Errorf("expected: true, actual: false")
	}
	if Same(tree.New(1), tree.New(2)) {
		t.Errorf("expected: false, actual: true")
	}
}
