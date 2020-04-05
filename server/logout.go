package server

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"jwt_server/jwt"
)

func logoutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	token, _ := getAccessToken(r)

	jwt.Delete(token)

	http.SetCookie(w, createAccessTokenCookie("", time.Now()))

	bytes, _ := ioutil.ReadFile("static/logout.html")
	w.Write(bytes)
}
