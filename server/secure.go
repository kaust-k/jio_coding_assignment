package server

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func securePathHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res, err := strconv.ParseBool(ps.ByName("authenticated"))
	loggedIn := res && err == nil

	val := struct {
		LoggedIn bool
	}{
		LoggedIn: loggedIn,
	}

	t, _ := template.ParseFiles("static/secure.tmpl")
	t.Execute(w, val)
}
