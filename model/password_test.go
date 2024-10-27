package model

import (
	"strings"
	"testing"
	"unicode"
)

// TestGeneratePasswords_MinMaxLength tests generation at minimum and maximum lengths.
func TestGeneratePasswords_MinMaxLength(t *testing.T) {
	// Test minimum length
	opts := PasswordOptions{
		Length:         6,
		Quantity:       3,
		IncludeSymbols: true,
		IncludeNumbers: true,
		IncludeUpper:   true,
		IncludeLower:   true,
	}
	passwords, err := GeneratePasswords(opts)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	for _, password := range passwords {
		if len(password) != 6 {
			t.Errorf("Expected password length of 6, but got %d", len(password))
		}
	}

	// Test maximum length
	opts.Length = 32
	passwords, err = GeneratePasswords(opts)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	for _, password := range passwords {
		if len(password) != 32 {
			t.Errorf("Expected password length of 32, but got %d", len(password))
		}
	}
}

// TestGeneratePasswords_OptionCombinations verifies behavior with various combinations of options.

func TestGeneratePasswords_OptionCombinations(t *testing.T) {
	combinations := []PasswordOptions{
		{Length: 10, IncludeSymbols: true, IncludeNumbers: false, IncludeUpper: false, IncludeLower: true},
		{Length: 10, IncludeSymbols: false, IncludeNumbers: true, IncludeUpper: true, IncludeLower: false},
		{Length: 10, IncludeSymbols: true, IncludeNumbers: true, IncludeUpper: true, IncludeLower: true},
		{Length: 10, IncludeSymbols: false, IncludeNumbers: false, IncludeUpper: true, IncludeLower: true},
	}

	for i, opts := range combinations {
		passwords, err := GeneratePasswords(opts)
		if err != nil {
			t.Errorf("Combination %d: Expected no error, but got %v", i+1, err)
			continue
		}

		for _, password := range passwords {
			if opts.IncludeSymbols {
				validateSymbols(t, password, i+1)
			}
			if opts.IncludeNumbers {
				validateNumbers(t, password, i+1)
			}
			if opts.IncludeUpper {
				validateUppercase(t, password, i+1)
			}
			if opts.IncludeLower {
				validateLowercase(t, password, i+1)
			}
		}
	}
}

// validateSymbols checks if symbols are present when IncludeSymbols is enabled.
func validateSymbols(t *testing.T, password string, comboNum int) {
	if !strings.ContainsAny(password, "!@#$%^&*()-_=+[]{}|;:,.<>/?") {
		t.Errorf("Combination %d: Expected symbols, but none found in password: %s", comboNum, password)
	}
}

// validateNumbers checks if numbers are present when IncludeNumbers is enabled.
func validateNumbers(t *testing.T, password string, comboNum int) {
	if !strings.ContainsAny(password, "0123456789") {
		t.Errorf("Combination %d: Expected numbers, but none found in password: %s", comboNum, password)
	}
}

// validateUppercase checks if uppercase letters are present when IncludeUpper is enabled.
func validateUppercase(t *testing.T, password string, comboNum int) {
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		t.Errorf("Combination %d: Expected uppercase letters, but none found in password: %s", comboNum, password)
	}
}

// validateLowercase checks if lowercase letters are present when IncludeLower is enabled.
func validateLowercase(t *testing.T, password string, comboNum int) {
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		t.Errorf("Combination %d: Expected lowercase letters, but none found in password: %s", comboNum, password)
	}
}

// TestGeneratePasswords_NoDuplicate tests generation with NoDuplicates option enabled.
func TestGeneratePasswords_NoDuplicate(t *testing.T) {
	opts := PasswordOptions{
		Length:       12,
		Quantity:     3,
		IncludeUpper: true,
		IncludeLower: true,
		NoDuplicates: true,
	}

	passwords, err := GeneratePasswords(opts)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	for _, password := range passwords {
		charCount := make(map[rune]int)
		for _, char := range password {
			charCount[char]++
			if charCount[char] > 1 {
				t.Errorf("Password %s contains duplicate characters", password)
				break
			}
		}
	}
}

// TestGeneratePasswords_NoSimilarCharacters tests generation with NoSimilar option enabled.
func TestGeneratePasswords_NoSimilarCharacters(t *testing.T) {
	opts := PasswordOptions{
		Length:       15,
		Quantity:     5,
		IncludeLower: true,
		NoSimilar:    true,
	}

	passwords, err := GeneratePasswords(opts)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	similarChars := "iIl1Lo0O"
	for _, password := range passwords {
		for _, char := range similarChars {
			if strings.ContainsRune(password, char) {
				t.Errorf("Password %s contains similar character %c", password, char)
			}
		}
	}
}

// TestGeneratePasswords_NoSequentialCharacters tests generation with NoSequential option enabled.
func TestGeneratePasswords_NoSequentialCharacters(t *testing.T) {
	opts := PasswordOptions{
		Length:         20,
		Quantity:       5,
		IncludeNumbers: true,
		IncludeLower:   true,
		NoSequential:   true,
	}

	passwords, err := GeneratePasswords(opts)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	for _, password := range passwords {
		runes := []rune(password)
		for i := 0; i < len(runes)-2; i++ {
			if isSequential(runes[i], runes[i+1], runes[i+2]) {
				t.Errorf("Password %s contains sequential characters", password)
			}
		}
	}
}

// TestPerformance_Consistency runs multiple iterations to ensure consistent behavior.
func TestPerformance_Consistency(t *testing.T) {
	opts := PasswordOptions{
		Length:         16,
		Quantity:       100,
		IncludeSymbols: true,
		IncludeNumbers: true,
		IncludeUpper:   true,
		IncludeLower:   true,
	}

	for i := 0; i < 10; i++ { // Run 10 iterations to check for consistency
		passwords, err := GeneratePasswords(opts)
		if err != nil {
			t.Errorf("Iteration %d: Expected no error, but got %v", i+1, err)
		}
		if len(passwords) != opts.Quantity {
			t.Errorf("Iteration %d: Expected %d passwords, but got %d", i+1, opts.Quantity, len(passwords))
		}
	}
}

// TestGetRandomLetter_OnlyLetters tests getRandomLetter to confirm it only generates letters.
func TestGetRandomLetter_OnlyLetters(t *testing.T) {
	opts := PasswordOptions{
		IncludeUpper: true,
		IncludeLower: true,
	}

	for i := 0; i < 20; i++ { // 20 iterations for better coverage
		char, err := getRandomLetter(opts)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
		if !unicode.IsLetter(rune(char)) {
			t.Errorf("Expected a letter, but got %c", char)
		}
	}
}
