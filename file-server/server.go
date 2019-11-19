package main

import (
	"./handlers"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func printHelp() {
	fmt.Println(`-name: user name
-pwd: user pwd
-port: 61000
		`)
}

func main() {
	printHelp()

	userName := flag.String("name", "admin", "-name: user name")
	userPwd := flag.String("pwd", "gato@123456", "-pwd: user pwd")
	port := flag.String("port", "61000", "-port: http port")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.FileHandler(w, r, userName, userPwd)
	})
	http.HandleFunc("/file/upload", func(w http.ResponseWriter, r *http.Request) {
		handlers.UploadHandler(w, r, userName, userPwd)
	})
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
