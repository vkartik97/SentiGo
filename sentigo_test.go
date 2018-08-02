package sentigo

import (
	"testing"
)

func TestPositive(t *testing.T) {
	actual := Score("This movie is good")
	if actual <= 0 {
		t.Fatalf("Failed")
	}

}

func TestNegative(t *testing.T) {
	actual := Score("This movie is bad")
	if actual >= 0 {
		t.Fatalf("Failed")
	}

}