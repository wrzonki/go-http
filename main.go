package main

import (
	"net/http"

	"github.com/wrzonki/http-server/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
