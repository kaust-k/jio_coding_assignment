package server

import (
	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

func init() {
	router = httprouter.New()

	router.GET("/login", getLoginHandler)
	router.POST("/login", postLoginHandler)

	router.GET("/logout", logoutHandler)

	router.GET("/secure", middleware(securePathHandler, authMiddleware))
}

func GetRouter() *httprouter.Router {
	return router
}
