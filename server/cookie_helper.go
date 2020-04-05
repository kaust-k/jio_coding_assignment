package server

import (
	"net/http"
	"time"
)

const (
	accessTokenCookieName = "access_token"
)

func getAccessTokenCookie(r *http.Request) (*http.Cookie, error) {
	return r.Cookie(accessTokenCookieName)
}

func createAccessTokenCookie(accessToken string, expiry time.Time) *http.Cookie {
	return &http.Cookie{
		Name:     accessTokenCookieName,
		Value:    accessToken,
		Expires:  expiry,
		HttpOnly: true,
	}
}
