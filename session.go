package main

import (
	"net/http"
	"github.com/gorilla/sessions"
	"errors"
)

var store = sessions.NewCookieStore([]byte("goblo-session"))

func currentUser(r *http.Request) (*Users, error) {
	session, err := store.Get(r, "goblo-session")
	if err != nil {
		return nil, err
	}

	uid := session.Values["uid"]

	if uid == nil {
		return nil, errors.New("Empty session")
	}

	user, err := findUser(int(uid.(int64)))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func createSession(user *Users, w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "goblo-session")
	if err != nil {
		return err
	}
	session.Values["uid"] = user.Id
	session.Save(r, w)
	return nil
}

func clearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "goblo-session")
	if err != nil {
		return err
	}
	delete(session.Values, "uid")
	session.Save(r, w)
	return nil
}

func setFlash(w http.ResponseWriter, r *http.Request, message string) error {
	session, err := store.Get(r, "goblo-session")
	if err != nil {
		return err
	}
	session.AddFlash(message)
	session.Save(r, w)
	return nil
}

func flashMessages(r *http.Request) ([]interface{}, error) {
	session, err := store.Get(r, "goblo-session")
	if err != nil {
		return nil, err
	}

	messages := session.Flashes()
	return messages, nil
}
