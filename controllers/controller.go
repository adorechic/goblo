package controllers

import (
	"errors"
	"github.com/adorechic/goblo/models"
	"github.com/gorilla/sessions"
	"github.com/russross/blackfriday"
	"html/template"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("goblo-session"))

type ViewObject struct {
	CurrentUser *models.User
	Pages       *[]models.Page
	Error       string
}

func render(w http.ResponseWriter, name string, data interface{}) {
	funcMap := template.FuncMap{
		"markdown": func(text string) template.HTML {
			markdown := blackfriday.MarkdownBasic([]byte(text))
			return template.HTML(string(markdown))
		},
	}
	t := template.New("")
	t = t.Funcs(funcMap)
	t = template.Must(t.ParseFiles("views/layout.html", "views/"+name+".html"))

	err := t.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		log.Println(err)
	}
}

func currentUser(r *http.Request) (*models.User, error) {
	session, err := store.Get(r, "goblo-session")
	if err != nil {
		return nil, err
	}

	uid := session.Values["uid"]

	if uid == nil {
		return nil, errors.New("Empty session")
	}

	user, err := models.FindUser(int(uid.(int64)))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func signin(w http.ResponseWriter, r *http.Request, username, password string) error {
	user, err := models.FindUserByCredential(username, password)
	if err != nil {
		return err
	}

	err = createSession(user, w, r)
	if err != nil {
		return err
	}

	return nil
}

func createSession(user *models.User, w http.ResponseWriter, r *http.Request) error {
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
