package server

import (
	"github.com/julienschmidt/httprouter"
)

func middleware(h httprouter.Handle, middleware ...func(httprouter.Handle) httprouter.Handle) httprouter.Handle {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}
