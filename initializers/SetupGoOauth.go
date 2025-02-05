package initializers

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var OauthConfig  *oauth2.Config


func InitOAuth() {
	ClientID     := os.Getenv("GOOGLE_CLIENT_ID")
	ClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	RedirectURL  := os.Getenv("BASE_INTERNAL_URL") + "/auth/google/callback"
	OauthConfig  = &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		RedirectURL:  RedirectURL,
		Scopes: []string{"email", "profile"},
		Endpoint: google.Endpoint,
	}
}
