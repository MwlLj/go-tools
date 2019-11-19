package handlers

import (
	"net/http"
)

func FileHandler(w http.ResponseWriter, r *http.Request, userName *string, userPwd *string) {
	authHandler(w, r, userName, userPwd, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path)
	})
}
