package initializers

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func SetupGoth() {
	// Set up the Google provider with the OAuth2 credentials
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:8080/auth/google/callback"),
	)

	// // Optional: You can log out or set up an error handler if desired
	// gothic.SetState("your-random-state-key") // Set a random state key for security
}
