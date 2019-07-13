package main

import (
	"./handler"
	"flag"
)

func main() {
	var host = flag.String("h", "localhost", "hostname of broker")
	var port = flag.String("p", "1883", "port of broker")
	var format = flag.String("f", handler.FormatJson, "format, if json -> -f json")
	var topic = flag.String("t", "#", "topic")
	var id = flag.String("id", "", "client id")
	var user = flag.String("user", "", "username")
	var pass = flag.String("pass", "", "password")
	var dump = flag.Bool("dump", false, "dump messages?")

	flag.Parse()

	sub := handler.NewSub()
	sub.Run(&handler.CRunParam{
		Host:     *host,
		Port:     *port,
		Format:   *format,
		Topic:    *topic,
		Id:       *id,
		UserName: *user,
		UserPwd:  *pass,
		Dump:     *dump,
	})
}
