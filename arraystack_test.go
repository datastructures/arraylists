package arraylists

import (
	"testing"
	"fmt"
)

func TestArrayStack(t *testing.T) {
	a := NewArrayStack()

	if a.Len() != 0 {
		t.Error("Size must be 0")
	}

	a.Push(1)
	a.Push(2)
	fmt.Println(a.array)
	if a.Len() != 2 {
		t.Error("Size must be 2")
	}

	a.Pop()
	fmt.Println(a.array)
	if a.Len() != 1 {
		t.Error("Size must be 1")
	}
}
