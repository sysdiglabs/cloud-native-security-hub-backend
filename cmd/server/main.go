package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sysdiglabs/prometheus-hub/web"
)

func main() {
	router := web.NewRouterWithLogger(log.New(os.Stderr, "", log.Ltime|log.Ldate|log.LUTC))

	log.Fatal(http.ListenAndServe(":8080", router))
}
