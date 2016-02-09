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

func clearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "goblo-session")
	if err != nil {
		return err
	}
	delete(session.Values, "uid")
	session.Save(r, w)
	return nil
}
