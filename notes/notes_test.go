// notes_test
package notes

import (
	"testing"
)

func TestGetTitle(t *testing.T) {
	n1 := NewNote("Hello World", 0)
	n2 := NewNote("Hello", 1)
	n3 := NewNote("", 2)
	n4 := NewNote("    ", 3)

	if n1.GetTitle() == "Hello" {
		t.Log("Success")
	} else {
		t.Fail()
	}

	if n2.GetTitle() == "Hello" {
		t.Log("Success")
	} else {
		t.Fail()
	}

	if n3.GetTitle() == "untitled" {
		t.Log("Success")
	} else {
		t.Logf("expected %v, got %v\n", "untitled", n4.GetTitle())
		t.Fail()
	}

	if n4.GetTitle() == "untitled" {
		t.Log("Success")
	} else {
		t.Logf("expected %v, got %v\n", "untitled", n4.GetTitle())
		t.Fail()
	}
}
