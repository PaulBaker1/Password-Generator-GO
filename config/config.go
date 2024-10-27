package config

import "awesomeProject/model"

// GetDefaultOptions initializes default password options.
func GetDefaultOptions() *model.PasswordOptions {
	return &model.PasswordOptions{
		MinLength:       6,
		MaxLength:       32,
		DefaultLength:   12,
		Quantity:        1,
		IncludeSymbols:  true,
		IncludeNumbers:  true,
		IncludeUpper:    true,
		IncludeLower:    true,
		BeginWithLetter: false,
		NoSimilar:       false,
		NoDuplicates:    false,
		NoSequential:    false,
	}
}
