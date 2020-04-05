package server

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"

	"jwt_server/jwt"
)

func getAccessToken(r *http.Request) (string, error) {
	cookie, err := getAccessTokenCookie(r)
	if err != nil {
		return "", err
	}
	if cookie.Value == "" {
		return "", errors.New("Access token not present")
	}

	return cookie.Value, nil
}

func verifyAccessToken(r *http.Request) (bool, error) {
	token, err := getAccessToken(r)
	if err == nil && token != "" {
		return jwt.Verify(token)
	}

	return false, err
}

func authMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ok, err := verifyAccessToken(r)
		verified := ok && err == nil
		ps = append(ps, httprouter.Param{Key: "authenticated", Value: strconv.FormatBool(verified)})
		next(w, r, ps)
	}
}
