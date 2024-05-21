package sessions

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"

	"github.com/joho/godotenv"
	"log"
	"os"
)

var store *sessions.CookieStore

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		log.Fatalf("SESSION_KEY not set in .env file")
	}

	store = sessions.NewCookieStore([]byte(sessionKey))
}

func GetSession(c echo.Context) (*sessions.Session, error) {
	return store.Get(c.Request(), "user-session")
}

func SetSessionValue(c echo.Context, key string, value interface{}) error {
	session, err := GetSession(c)
	if err != nil {
		return err
	}

	session.Values[key] = value
	return session.Save(c.Request(), c.Response())
}

func GetSessionValue(c echo.Context, key string) (interface{}, error) {
	session, err := GetSession(c)
	if err != nil {
		return nil, err
	}

	return session.Values[key], nil
}
