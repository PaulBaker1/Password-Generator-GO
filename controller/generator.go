/**
 * Password Generator - Controller
 *
 * This file contains the GeneratorController struct and methods for managing
 * password generation. The controller interacts with the model to generate
 * passwords based on user-defined options passed from the GUI.
 */

package controller

import (
	"awesomeProject/config"
	"awesomeProject/model"
)

// GeneratorController manages password generation requests.
// Purpose:
//
//	Manages and coordinates password generation requests from the view by
//	interfacing with the password generation logic in the model.
type GeneratorController struct {
	Config *model.PasswordOptions
}

// NewGeneratorController initializes the controller with default options.
// Purpose:
//
//	Create a new instance of the GeneratorController with default configurations.
//
// Returns:
//
//	*GeneratorController: An instance of the GeneratorController with default settings.
//
// Example:
//
//	ctrl := NewGeneratorController()
func NewGeneratorController() *GeneratorController {
	return &GeneratorController{
		Config: config.GetDefaultOptions(),
	}
}

// GeneratePasswords generates a list of passwords based on the options provided.
// Parameters:
//   - opts (model.PasswordOptions): The settings used to customize password generation.
//
// Returns:
//
//	[]string: A list of generated passwords based on the quantity specified in opts.
//	error: Returns an error if password generation fails due to invalid options.
//
// Example:
//
//	passwords, err := ctrl.GeneratePasswords(opts)
func (gc *GeneratorController) GeneratePasswords(opts model.PasswordOptions) ([]string, error) {
	return model.GeneratePasswords(opts)
}
