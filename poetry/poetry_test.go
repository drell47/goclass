package poetry

import (
	"testing"
)

var p1 Poem
var p0 Poem

func init() {
	p1 = Poem{{"Line 1 adg, frog?", "Line 2 up/down, uiop.", "Line 3 trewgfd,", "Line 4, bvcxde."}}
	p0 = Poem{}
}

func TestNumWords(t *testing.T) {
	if p0.NumWords() != 0 {
		t.Fatalf("Unexpected word count - expected 0, got %d\n", p0.NumWords())
	}
	if p1.NumWords() != 14 {
		t.Fatalf("Unexpected word count - expected 14, got %d\n", p1.NumWords())
	}
}

func TestNumStanzas(t *testing.T) {

	if p0.NumStanzas() != 0 {
		t.Fatalf("Unexpected stanza count - expected 0, got %d", p0.NumStanzas())
	}
	if p1.NumStanzas() != 1 {
		t.Fatalf("Unexpected stanza count - expected 1, got %d", p1.NumStanzas())
	}

	if p0.NumLines() != 0 {
		t.Fatalf("Unexpected line count - expected 0, got %d", p0.NumLines())
	}
	if p1.NumLines() != 4 {
		t.Fatalf("Unexpected line count - expected 4, got %d", p1.NumLines())
	}
}

func TestChars(t *testing.T) {

	v, c, punct := p0.Stats()
	if v != 0 || c != 0 || punct != 0 {
		t.Fatalf("wrong numbers")
	}

	v, c, punct = p1.Stats()
	if v != 17 || c != 33 || punct != 18 {
		t.Fatalf("wrong v: %d, c: %d, punct: %d\n", v, c, punct)
	}
}
