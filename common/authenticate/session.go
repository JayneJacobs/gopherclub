package authenticate

import (
	"log"
	"net/http"
	"os"

	"github.com/JayneJacobs/gopherclub/models"

	"github.com/gorilla/sessions"
)

var SessionStore *sessions.FilesystemStore

func CreateUserSession(u *models.User, sessionID string, w http.ResponseWriter, r *http.Request) error {

	gfSession, err := SessionStore.Get(r, "gopherclub-session")

	if err != nil {
		log.Print(err)
	}

	gfSession.Values["sessionID"] = sessionID
	gfSession.Values["username"] = u.Username
	gfSession.Values["firstName"] = u.FirstName
	gfSession.Values["lastName"] = u.LastName
	gfSession.Values["email"] = u.Email
	gfSession.Values["uuid"] = u.UUID

	err = gfSession.Save(r, w)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func ExpireUserSession(w http.ResponseWriter, r *http.Request) {
	gfSession, err := SessionStore.Get(r, "gopherclub-session")

	if err != nil {
		log.Print(err)
	}

	gfSession.Options.MaxAge = -1
	gfSession.Save(r, w)
}

func init() {
	if _, err := os.Stat("/tmp/gopherclub-sessions"); os.IsNotExist(err) {
		os.Mkdir("/tmp/gopherclub-sessions", 711)
	}
	SessionStore = sessions.NewFilesystemStore("/tmp/gopherclub-sessions", []byte(os.Getenv("gopherclub_HASH_KEY")))
}
