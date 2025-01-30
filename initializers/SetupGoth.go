package initializers

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/linkedin"
)

func SetupGoth() {
	// Set up the Google provider with the OAuth2 credentials
	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),          // the client ID
			os.Getenv("GOOGLE_CLIENT_SECRET"),      // the client secret
			os.Getenv("BASE_INTERNAL_URL") + "/auth/google/callback", // Callback URL
			"https://www.googleapis.com/auth/userinfo.email", // Valid scope
			"https://www.googleapis.com/auth/userinfo.profile", // Valid scope
		),
		linkedin.New(
			os.Getenv("LINKEDIN_CLIENT_ID"), 		// the client ID
			os.Getenv("LINKEDIN_CLIENT_SECRET"),	// the client secret
			os.Getenv("BASE_INTERNAL_URL") + "/auth/linkedin/callback",		// Callback URL
			"r_liteprofile", "r_emailaddress", // Valid scope
		),
	)

	// // Optional: You can log out or set up an error handler if desired
	// gothic.SetState("your-random-state-key") // Set a random state key for security
}
