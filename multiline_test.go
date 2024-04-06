package multiline

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	actual := String(`Line 1
    Line 2
  Line 3
         Line 4
    `)

	expected := `Line 1
Line 2
Line 3
Line 4
    `
	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("expected >>%v<< but got >>%v<<", expected, actual)
	}

}
