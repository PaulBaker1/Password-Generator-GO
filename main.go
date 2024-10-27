/**
 * Password Generator - Main Entry Point
 *
 * This file serves as the entry point for the password generator application,
 * initializing the controller and launching the GUI. The main function
 * sets up the default configurations and triggers the GUI layout.
 */

package main

import (
	"password-generator/controller"
	"password-generator/view"
)

// main initializes the password generator's controller and launches the GUI.
// Purpose:
//
//	Set up the password generator's configurations and start the application GUI.
//
// Example:
//
//	Run the main function to start the application: go run main.go
func main() {
	// Initialize the controller with default options
	ctrl := controller.NewGeneratorController()

	// Start the GUI and pass the controller
	view.StartGUI(ctrl)
}
