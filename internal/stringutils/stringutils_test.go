package stringutils_test

import (
	"passfish/internal/stringutils"
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	// kitten -> mitten (1 substitution)
	if stringutils.LevenshteinDistance("kitten", "mitten") != 1 {
		t.Error("Expected 1, got something else")
	}

	// kitten -> sitting (3 substitutions)
	if stringutils.LevenshteinDistance("kitten", "sitting") != 3 {
		t.Error("Expected 1, got something else")
	}
}

func TestTopMatches(t *testing.T) {
	words := []string{"mitten", "sitting", "barf"}
	sol := stringutils.FindTopMatches(words, "kitten", 1)
	if len(sol) != 1 {
		t.Errorf("Expected len(1), got len(%d)", len(sol))
	}
	if sol[0] != "mitten" {
		t.Errorf("Expected 'mitten', got '%s'", sol[0])
	}

	sol = stringutils.FindTopMatches(words, "kitten", 2)
	if len(sol) != 2 {
		t.Errorf("Expected len(2), got len(%d)", len(sol))
	}
	if sol[0] != "mitten" {
		t.Errorf("Expected 'mitten', got '%s'", sol[0])
	}
	if sol[1] != "sitting" {
		t.Errorf("Expected 'mitten', got '%s'", sol[1])
	}

	sol = stringutils.FindTopMatches(words, "kitten", 3)
	if len(sol) != 3 {
		t.Errorf("Expected len(3), got len(%d)", len(sol))
	}
	if sol[0] != "mitten" {
		t.Errorf("Expected 'mitten', got '%s'", sol[0])
	}
	if sol[1] != "sitting" {
		t.Errorf("Expected 'sitting', got '%s'", sol[1])
	}
	if sol[2] != "barf" {
		t.Errorf("Expected 'barf', got '%s'", sol[2])
	}

	// try another ordering
	words = []string{"barf", "mitten", "sitting"}
	sol = stringutils.FindTopMatches(words, "kitten", 1)
	if len(sol) != 1 {
		t.Errorf("Expected len(1), got len(%d)", len(sol))
	}
	if sol[0] != "mitten" {
		t.Errorf("Expected 'mitten', got '%s'", sol[0])
	}

	sol = stringutils.FindTopMatches(words, "kitten", 2)
	if len(sol) != 2 {
		t.Errorf("Expected len(2), got len(%d)", len(sol))
	}
	if sol[0] != "mitten" {
		t.Errorf("Expected 'mitten', got '%s'", sol[0])
	}
	if sol[1] != "sitting" {
		t.Errorf("Expected 'sitting', got '%s'", sol[1])
	}

	sol = stringutils.FindTopMatches(words, "kitten", 3)
	if len(sol) != 3 {
		t.Errorf("Expected len(3), got len(%d)", len(sol))
	}
	if sol[0] != "mitten" {
		t.Errorf("Expected 'mitten', got '%s'", sol[0])
	}
	if sol[1] != "sitting" {
		t.Errorf("Expected 'sitting', got '%s'", sol[1])
	}
	if sol[2] != "barf" {
		t.Errorf("Expected 'barf', got '%s'", sol[2])
	}

	// try less words than requested
	words = []string{"barf", "mitten", "sitting"}
	sol = stringutils.FindTopMatches(words, "kitten", 5)
	if len(sol) != 3 {
		t.Errorf("Expected len(3), got len(%d)", len(sol))
	}
	if sol[0] != "mitten" {
		t.Errorf("Expected 'mitten', got '%s'", sol[0])
	}
	if sol[1] != "sitting" {
		t.Errorf("Expected 'sitting', got '%s'", sol[1])
	}
	if sol[2] != "barf" {
		t.Errorf("Expected 'barf', got '%s'", sol[2])
	}

}
