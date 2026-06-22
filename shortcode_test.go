package shortcode

import (
	"slices"
	"testing"
)

func TestLookup(t *testing.T) {
	em, ok := Lookup("smile")
	if !ok {
		t.Fatal("expected smile to be found")
	}
	if em != "😄" {
		t.Fatalf("got %q want 😄", em)
	}

	if _, ok := Lookup("not_a_real_emoji_12345"); ok {
		t.Fatal("expected unknown shortcode to be missing")
	}
}

func TestLookupWithColons(t *testing.T) {
	em, ok := Lookup(":wave:")
	if !ok {
		t.Fatal("expected :wave: to be found")
	}
	if em != "👋" {
		t.Fatalf("got %q want 👋", em)
	}
}

func TestLookupCaseInsensitive(t *testing.T) {
	em, ok := Lookup("Smile")
	if !ok || em != "😄" {
		t.Fatalf("got %q, ok=%v", em, ok)
	}
}

func TestSuggest(t *testing.T) {
	matches := Suggest("sm")
	if len(matches) == 0 {
		t.Fatal("expected matches for sm")
	}
	if matches[0].Code != "smile" {
		t.Fatalf("expected shortest match first, got %q", matches[0].Code)
	}
	foundSmile := false
	for _, m := range matches {
		if m.Code == "smile" {
			foundSmile = true
		}
	}
	if !foundSmile {
		t.Fatal("expected smile in suggestions")
	}
}

func TestSuggestEmpty(t *testing.T) {
	if Suggest("") != nil {
		t.Fatal("empty prefix should return nil")
	}
	if Suggest("   ") != nil {
		t.Fatal("whitespace-only prefix should return nil")
	}
}

func TestAllSorted(t *testing.T) {
	all := All()
	if len(all) == 0 {
		t.Fatal("expected some entries")
	}
	codes := make([]string, len(all))
	for i, m := range all {
		codes[i] = m.Code
	}
	if !slices.IsSorted(codes) {
		t.Fatal("All() must be sorted by shortcode")
	}
}

func TestAllEntriesUnique(t *testing.T) {
	all := All()
	seen := make(map[string]struct{}, len(all))
	for _, m := range all {
		if _, ok := seen[m.Code]; ok {
			t.Fatalf("duplicate shortcode: %s", m.Code)
		}
		seen[m.Code] = struct{}{}
	}
}
