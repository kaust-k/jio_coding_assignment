package server

import (
	"html/template"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"jwt_server/jwt"
)

func getLoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("static/login.html")
	t.Execute(w, nil)
}

func postLoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if ok, _ := verifyAccessToken(r); ok {
		w.WriteHeader(http.StatusOK)
		return
	}

	username, password, ok := r.BasicAuth()

	if !ok || !validateCredentials(username, password) {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	token, expiry, err := jwt.Create(username)
	if err == nil {
		http.SetCookie(w, createAccessTokenCookie(token, expiry))
		io.WriteString(w, `{"status":"ok"}`)
	}
}

func validateCredentials(username, password string) bool {
	return username == "test" && password == "test"
}
