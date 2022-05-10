package main

import (
	"example/crud/router"
	"net/http"
)

func main() {
	// Server listening on localhost:
	http.ListenAndServe(router.HttpAddr, router.Start())
}
