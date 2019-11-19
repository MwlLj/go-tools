package handlers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func authHandler(w http.ResponseWriter, r *http.Request, userName *string, userPwd *string, success func(w http.ResponseWriter, r *http.Request)) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println(auth)
	auths := strings.SplitN(auth, " ", 2)
	if len(auths) != 2 {
		fmt.Println("error")
		return
	}
	authMethod := auths[0]
	authB64 := auths[1]
	switch authMethod {
	case "Basic":
		authstr, err := base64.StdEncoding.DecodeString(authB64)
		if err != nil {
			fmt.Println(err)
			io.WriteString(w, "Unauthorized!\n")
			return
		}
		fmt.Println(string(authstr))
		pwd := strings.SplitN(string(authstr), ":", 2)
		if len(pwd) != 2 {
			fmt.Println("error")
			io.WriteString(w, "error!\n")
			return
		}
		username := pwd[0]
		password := pwd[1]
		if username != *userName || password != *userPwd {
			fmt.Println("name or pwd error")
			io.WriteString(w, "name or pwd error!\n")
			return
		}
	default:
		fmt.Println("error")
		return
	}
	success(w, r)
}
