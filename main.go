package main

import (
	"net/http"

	"github.com/gloryof/go-crud-practice/context/user/web"
)

func main() {
	server := http.Server{
		Addr: ":8000",
	}

	welcomeHandler := handlers.WelcomeHandler{}

	http.Handle("/welcome", &welcomeHandler)

	server.ListenAndServe()
}
