package initializers

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func SetupGoth() {
	// Set up the Google provider with the OAuth2 credentials
	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),          // Correctly set the client ID
			os.Getenv("GOOGLE_CLIENT_SECRET"),      // Correctly set the client secret
			os.Getenv("BASE_INTERNAL_URL") + "/auth/google/callback", // Redirect URI
			"https://www.googleapis.com/auth/userinfo.email", // Valid scope
			"https://www.googleapis.com/auth/userinfo.profile", // Valid scope
		),
	)

	// // Optional: You can log out or set up an error handler if desired
	// gothic.SetState("your-random-state-key") // Set a random state key for security
}
