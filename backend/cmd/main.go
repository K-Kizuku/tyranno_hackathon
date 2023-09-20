package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"tyranno/backend/adapter"
)

func main() {
	var (
		port = flag.String("port", "8080", "addr to bind")
	)
	flag.Parse()
	s := adapter.New()
	s.Init()
	s.Middleware()
	s.InitRouter()
	log.Println(http.ListenAndServe(fmt.Sprint(":", *port), s.Router))
}
