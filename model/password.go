/**
 * Password Generation Logic
 *
 * This file contains the core password generation logic, structured around the
 * PasswordOptions struct, which holds all customizable parameters. The functions
 * here generate secure passwords based on user-defined options, ensuring that
 * specific criteria are met, such as the exclusion of similar, duplicate, or
 * sequential characters.
 */

package model

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

// PasswordOptions holds user-selected settings for password customization.
// Purpose:
//
//	Used to define various aspects of password generation, such as the inclusion
//	of symbols, length of the password, and the exclusion of certain characters.
//
// Fields:
//   - MinLength, MaxLength, DefaultLength (int): Controls length limits and defaults.
//   - Quantity (int): Specifies the number of passwords to generate.
//   - IncludeSymbols, IncludeNumbers, IncludeUpper, IncludeLower (bool): Defines
//     character types to include.
//   - BeginWithLetter, NoSimilar, NoDuplicates, NoSequential (bool): Additional
//     customization options for password structure.
//   - Length (int): Desired length for each password.
type PasswordOptions struct {
	MinLength       int
	MaxLength       int
	DefaultLength   int
	Quantity        int
	IncludeSymbols  bool
	IncludeNumbers  bool
	IncludeUpper    bool
	IncludeLower    bool
	BeginWithLetter bool
	NoSimilar       bool
	NoDuplicates    bool
	NoSequential    bool
	Length          int
}

// similarCharacters holds a string of visually similar characters
// (e.g., "i" and "l") that should be excluded from generated passwords if
// the NoSimilar option is enabled.
var similarCharacters = "iIl1Lo0O"

// GeneratePasswords generates a list of passwords based on the provided options.
// Purpose:
//
//	Generates multiple passwords using the GeneratePassword function, iterating
//	based on the Quantity field in PasswordOptions.
//
// Parameters:
//   - opts (PasswordOptions): Settings used to customize the passwords generated.
//
// Returns:
//
//	[]string: A list of generated passwords.
//	error: Returns an error if password generation fails due to invalid options.
//
// Example:
//
//	passwords, err := GeneratePasswords(opts)
func GeneratePasswords(opts PasswordOptions) ([]string, error) {
	var passwords []string
	for i := 0; i < opts.Quantity; i++ {
		password, err := generatePassword(opts)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, password)
	}
	return passwords, nil
}

// generatePassword creates a single password based on the options provided.
// Purpose:
//
//	Builds a password character set and assembles the password according to user
//	specifications, ensuring that specific structural requirements are met.
//
// Parameters:
//   - opts (PasswordOptions): Settings for password length, character types, and restrictions.
//
// Returns:
//
//	string: A generated password.
//	error: An error if no valid character types are selected.
//
// Example:
//
//	password, err := generatePassword(opts)
func generatePassword(opts PasswordOptions) (string, error) {
	chars := buildCharacterSet(opts)
	if chars == "" {
		return "", errors.New("at least one character type must be selected")
	}

	password := make([]byte, opts.Length)
	var err error

	for i := 0; i < opts.Length; i++ {
		if i == 0 && opts.BeginWithLetter {
			password[i], err = getRandomLetter(opts)
		} else {
			password[i], err = secureRandomChar(chars)
		}
		if err != nil {
			return "", err
		}
	}

	passwordStr := string(password)

	// Post-process to ensure no similar, duplicate, or sequential characters
	if opts.NoSimilar {
		passwordStr = removeSimilarCharacters(passwordStr)
	}
	if opts.NoDuplicates {
		passwordStr = removeDuplicateCharacters(passwordStr)
	}
	if opts.NoSequential {
		passwordStr = removeSequentialCharacters(passwordStr)
	}

	return passwordStr, nil
}

// buildCharacterSet compiles a set of allowed characters based on options.
// Purpose:
//
//	Builds a character set according to user-specified options for symbols,
//	numbers, uppercase, and lowercase letters.
//
// Parameters:
//   - opts (PasswordOptions): Settings that determine the characters included.
//
// Returns:
//
//	string: A string containing the allowed characters for password generation.
func buildCharacterSet(opts PasswordOptions) string {
	var chars string
	if opts.IncludeSymbols {
		chars += "!@#$%^&*()-_=+[]{}|;:,.<>/?"
	}
	if opts.IncludeNumbers {
		chars += "0123456789"
	}
	if opts.IncludeUpper {
		chars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if opts.IncludeLower {
		chars += "abcdefghijklmnopqrstuvwxyz"
	}
	return chars
}

// getRandomLetter retrieves a random letter from the allowed letter set.
// Purpose:
//
//	Selects a random letter (uppercase or lowercase) when passwords must begin
//	with a letter or to comply with the user’s selection criteria.
//
// Parameters:
//   - opts (PasswordOptions): Specifies whether uppercase or lowercase letters are allowed.
//
// Returns:
//
//	byte: A randomly selected letter from the allowed set.
//	error: An error if no valid letter options are available.
func getRandomLetter(opts PasswordOptions) (byte, error) {
	letters := ""
	if opts.IncludeUpper {
		letters += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if opts.IncludeLower {
		letters += "abcdefghijklmnopqrstuvwxyz"
	}
	return secureRandomChar(letters)
}

// secureRandomChar returns a random character from a given character set.
// Purpose:
//
//	Provides cryptographic security for character selection using crypto/rand.
//
// Parameters:
//   - chars (string): The set of characters to choose from.
//
// Returns:
//
//	byte: A securely generated random character.
//	error: An error if secure random generation fails.
func secureRandomChar(chars string) (byte, error) {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
	if err != nil {
		return 0, errors.New("failed to generate secure random character")
	}
	return chars[index.Int64()], nil
}

// removeSimilarCharacters removes visually similar characters from the password.
// Purpose:
//
//	Enhances readability by removing similar characters if NoSimilar is enabled.
//
// Parameters:
//   - password (string): The original password string.
//
// Returns:
//
//	string: The password string with similar characters removed.
func removeSimilarCharacters(password string) string {
	for _, char := range similarCharacters {
		password = strings.ReplaceAll(password, string(char), "")
	}
	return password
}

// removeDuplicateCharacters removes duplicate characters from the password.
// Purpose:
//
//	Ensures each character appears only once if NoDuplicates is enabled.
//
// Parameters:
//   - password (string): The original password string.
//
// Returns:
//
//	string: The password string with duplicate characters removed.
func removeDuplicateCharacters(password string) string {
	seen := make(map[rune]bool)
	result := strings.Builder{}
	for _, char := range password {
		if !seen[char] {
			seen[char] = true
			result.WriteRune(char)
		}
	}
	return result.String()
}

// removeSequentialCharacters detects and replaces sequential characters in the password.
// Purpose:
//
//	Prevents the use of ascending or descending sequences if NoSequential is enabled.
//
// Parameters:
//   - password (string): The original password string.
//
// Returns:
//
//	string: The password with sequential characters replaced.
func removeSequentialCharacters(password string) string {
	var result strings.Builder
	runes := []rune(password)

	for i := 0; i < len(runes); i++ {
		if i+2 < len(runes) && isSequential(runes[i], runes[i+1], runes[i+2]) {
			// Replace the sequence with random non-sequential characters
			replacement := generateNonSequentialChars(runes, i)
			result.WriteString(replacement)
			i += 2 // Skip the next two characters as they are part of the sequence
		} else {
			result.WriteRune(runes[i])
		}
	}

	return result.String()
}

// isSequential checks if three characters form a sequence.
// Purpose:
//
//	Determines if three characters form an ascending or descending sequence.
//
// Parameters:
//   - a, b, c (rune): Three consecutive characters from the password.
//
// Returns:
//
//	bool: True if the characters are sequential; otherwise, false.
func isSequential(a, b, c rune) bool {
	return (b == a+1 && c == b+1) || (b == a-1 && c == b-1)
}

// generateNonSequentialChars generates three random characters that are non-sequential.
// Purpose:
//
//	Replaces sequential characters with random non-sequential characters.
//
// Parameters:
//   - runes ([]rune): The password characters.
//   - index (int): The index of the sequence start.
//
// Returns:
//
//	string: A string of non-sequential characters to replace the sequence.
func generateNonSequentialChars(runes []rune, index int) string {
	var replacementRunes []rune
	for len(replacementRunes) < 3 {
		randomChar := getRandomRune()
		if (index > 0 && isSequential(runes[index-1], randomChar, ' ')) ||
			(index+3 < len(runes) && isSequential(randomChar, runes[index+3], ' ')) {
			continue // Skip this character if it forms a sequence
		}
		replacementRunes = append(replacementRunes, randomChar)
	}
	return string(replacementRunes)
}

// getRandomRune generates a random rune from a preset character set.
// Purpose:
//
//	Used to obtain a random character that does not introduce sequential patterns.
//
// Returns:
//
//	rune: A randomly selected character from the character set.
func getRandomRune() rune {
	charSets := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charSets))))
	return rune(charSets[index.Int64()])
}
