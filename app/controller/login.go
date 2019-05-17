package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Momper14/web/app/model"
	"github.com/Momper14/weblib/client"
	"github.com/Momper14/weblib/client/users"
	"github.com/gorilla/sessions"
)

const (
	// SessionCookieName name of the session cookie
	SessionCookieName = "sessionid"
)

var (
	key   = []byte{239, 83, 234, 126, 210, 130, 176, 193, 49, 189, 102, 247, 235, 72, 30, 96, 109, 235, 12, 47, 86, 171, 20, 133, 65, 40, 180, 152, 242, 245, 150, 232}
	store = sessions.NewCookieStore(key)
)

// LoginController controller for login
func LoginController(w http.ResponseWriter, r *http.Request) {
	defer recoverInternalError()

	session, err := store.Get(r, SessionCookieName)
	if err != nil {
		internalError(err, w)
	}

	if !IstEingeloggt(w, r) {
		decoder := json.NewDecoder(r.Body)
		var login model.Login

		if err := decoder.Decode(&login); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			panic(err)
		}

		if l := len(login.Passwort); l != 128 {
			http.Error(w, fmt.Sprintf("Password: %s hat eine ungültige SHA-512 Zeichenlänge von %d", login.Passwort, l), http.StatusBadRequest)
			return
		}

		user, err := users.New().UserByID(login.User)
		if err != nil {
			if _, ok := err.(client.NotFoundError); ok {
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintln(w, "User oder Password falsch")
				return
			}
			internalError(err, w)
		}

		if user.Password != login.Passwort {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, "User oder Password falsch")
			return
		}

		session.Values["authenticated"] = true
		session.Values["user"] = login.User
		session.Save(r, w)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Logged in")
	} else {
		w.WriteHeader(http.StatusAlreadyReported)
		fmt.Fprintln(w, "already logged in")
	}
}

// LogoutController controller for logout
func LogoutController(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, SessionCookieName)
	if err != nil {
		internalError(err, w)
	}
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		delete(session.Values, "authenticated")
		delete(session.Values, "user")
		session.Save(r, w)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Logged out")
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, "You are not logged in")
}

// IstEingeloggt checks if session is logged in
func IstEingeloggt(w http.ResponseWriter, r *http.Request) bool {
	session, err := store.Get(r, SessionCookieName)
	if err != nil {
		internalError(err, w)
	}
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		return true
	}
	return false
}

// GetUser Returns the current User
func GetUser(w http.ResponseWriter, r *http.Request) string {
	session, err := store.Get(r, SessionCookieName)
	if err != nil {
		internalError(err, w)
	}

	if user, ok := session.Values["user"].(string); ok {
		return user
	}

	internalError(fmt.Errorf("Error: Kein user vorhanden"), w)
	return ""
}
