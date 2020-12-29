package main

import (
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (sess data.Sesstion, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Sesstion{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.new("Invalid session")
		}
	}
	return
}
