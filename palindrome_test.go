package bluebird

import "testing"

func TestCleanString(t *testing.T) {
	expected := "malammalam"
	if observed := CleanString("Malam? malaM"); observed != expected {
		t.Fatalf("CleanString() = %v, want %v", observed, expected)
	}
}

func TestValidatePalindrome(t *testing.T) {
	longStr := "Was it a car or a cat I saw?"
	expected := true
	if isPalindrome := ValidatePalindrome("Malam? malaM"); isPalindrome != expected {
		t.Fatalf("ValidatePalindrome() = %v, want %v", isPalindrome, expected)
	}

	if isPalindrome := ValidatePalindrome(longStr); isPalindrome != expected {
		t.Fatalf("ValidatePalindrome() = %v, want %v", isPalindrome, expected)
	}
}
