package main

import "fmt"

// On the first application run, GoLand will ask if you want to
// keep sudo running or terminate it immediately. If you would
// like to change your response, go to
// Preferences/Settings | Appearance & Behavior | System Settings | Process Elevation.
//
// Note: By enabling this option you grant GoLand and all third-party
// plugins access to your system.

// Step 1. Run the application and observe you are (likely) not an admin/root user.
// Step 2. Click on the run button and select Modify Run Configuration...
// Step 3. Enable the "Run with elevated privileges"/"Run with sudo" for your configuration.
// Step 4. Run the application again.

func main() {
	fmt.Printf("Is running as admin? %t\n", isElevated())
}
