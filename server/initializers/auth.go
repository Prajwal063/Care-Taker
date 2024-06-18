package initializers

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	MaxAge = 86400
	IsProd = false
)

func AuthStore() cookie.Store {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	store := cookie.NewStore([]byte(os.Getenv("AUTH_SECRET")))
	store.Options(sessions.Options{
		MaxAge:   MaxAge,
		Path:     "/",
		HttpOnly: true,
		Secure:   IsProd,
	})

	gothic.Store = store

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	goth.UseProviders(
		google.New(googleClientID, googleSecret, "http://localhost:8000/auth/google/callback", "email", "profile"),
	)

	return store
}
