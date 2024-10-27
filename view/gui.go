/**
 * Password Generator - GUI Layout
 *
 * This file defines the graphical user interface for the password generator
 * using the Fyne framework. It includes the layout configuration, GUI components,
 * and interactions for password generation and display.
 */

package view

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"awesomeProject/controller"
	"awesomeProject/model"
)

// StartGUI initializes and runs the GUI layout for the password generator.
// Purpose:
//
//	Sets up the GUI layout and components for the password generator using Fyne.
//	Allows users to select password generation options and displays generated passwords.
//
// Parameters:
//   - ctrl (*controller.GeneratorController): The controller that manages password generation.
//
// Example:
//
//	StartGUI(ctrl)
func StartGUI(ctrl *controller.GeneratorController) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Password Generator")

	// Set up the length slider with min, max, and default values from the controller config
	lengthSlider := widget.NewSlider(float64(ctrl.Config.MinLength), float64(ctrl.Config.MaxLength))
	lengthSlider.Value = float64(ctrl.Config.DefaultLength)
	lengthLabel := widget.NewLabel(fmt.Sprintf("Length: %.0f", lengthSlider.Value))
	lengthSlider.OnChanged = func(value float64) {
		lengthLabel.SetText(fmt.Sprintf("Length: %.0f", value))
	}

	// Quantity selection dropdown to determine how many passwords to generate.
	quantitySelect := widget.NewSelect([]string{"1", "5", "10", "20"}, nil)
	quantitySelect.SetSelected("1") // Default selection to 1 password

	// Options for character inclusion in the generated password
	includeSymbols := widget.NewCheck("Include Symbols", nil)
	includeNumbers := widget.NewCheck("Include Numbers", nil)
	includeUpper := widget.NewCheck("Include Uppercase Letters", nil)
	includeLower := widget.NewCheck("Include Lowercase Letters", nil)
	includeLower.SetChecked(true) // Default to lowercase inclusion

	// Additional options for password customization
	beginWithLetter := widget.NewCheck("Begin With Letters", nil)
	noSimilar := widget.NewCheck("No Similar Characters", nil)
	noDuplicates := widget.NewCheck("No Duplicate Characters", nil)
	noSequential := widget.NewCheck("No Sequential Characters", nil)

	// passwordEntry allows generated passwords to be displayed and edited.
	passwordEntry := widget.NewMultiLineEntry()
	passwordEntry.SetPlaceHolder("Generated passwords will appear here")
	passwordEntry.Wrapping = fyne.TextWrapWord // Allows word wrapping for multi-line display

	// Generate Button
	// Purpose: Triggers password generation based on selected options.
	// Example:
	//   Clicking the button generates and displays passwords.
	generateButton := widget.NewButton("Generate", func() {
		// Convert selected quantity to integer
		quantity, err := strconv.Atoi(quantitySelect.Selected)
		if err != nil {
			passwordEntry.SetText("Error: invalid quantity selected")
			return
		}

		// Set up password options for generation
		opts := model.PasswordOptions{
			Length:          int(lengthSlider.Value),
			Quantity:        quantity,
			IncludeSymbols:  includeSymbols.Checked,
			IncludeNumbers:  includeNumbers.Checked,
			IncludeUpper:    includeUpper.Checked,
			IncludeLower:    includeLower.Checked,
			BeginWithLetter: beginWithLetter.Checked,
			NoSimilar:       noSimilar.Checked,
			NoDuplicates:    noDuplicates.Checked,
			NoSequential:    noSequential.Checked,
		}

		// Generate passwords and display them in a numbered format
		passwords, err := ctrl.GeneratePasswords(opts)
		if err != nil {
			passwordEntry.SetText("Error: " + err.Error())
		} else {
			var formattedPasswords strings.Builder
			for i, password := range passwords {
				formattedPasswords.WriteString(fmt.Sprintf("%d. %s\n", i+1, password))
			}
			passwordEntry.SetText(formattedPasswords.String())
		}
	})

	// Layout configuration - passwordEntry expands to fill available space.
	content := container.NewBorder(
		container.NewVBox(
			widget.NewLabel("Password Generator"),
			lengthLabel,
			lengthSlider,
			quantitySelect,
			includeSymbols,
			includeNumbers,
			includeUpper,
			includeLower,
			beginWithLetter,
			noSimilar,
			noDuplicates,
			noSequential,
			generateButton,
		),
		nil, nil, nil, passwordEntry, // passwordEntry fills remaining space
	)

	// Set the content and display the window
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 500)) // Initial window size
	myWindow.ShowAndRun()
}
