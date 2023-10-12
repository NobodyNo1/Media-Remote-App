package main

import (
	"log"
	"net/http"
	"remote_media_control/router"
)

func main() {
	// start the server
	router.SetupRouter()
	log.Default().Print(http.ListenAndServe(":8000", nil))
}
