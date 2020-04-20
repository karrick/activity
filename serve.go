package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/karrick/gohm/v2"
	"github.com/karrick/golf"
)

func main() {
	optPort := golf.IntP('p', "port", 8080, "HTTP server network port")
	optStatic := golf.StringP('s', "static", "static", "filesystem pathname to static virtual root")
	golf.Parse()

	*optStatic = filepath.Clean(*optStatic)

	mux := http.NewServeMux()

	mux.Handle("/static/", gohm.StaticHandler("/static", *optStatic))

	mux.Handle("/", gohm.DefaultHandler(filepath.Join(*optStatic, "index.html")))

	log.Print("[INFO] web server port: ", *optPort)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *optPort),
		Handler: gohm.New(gohm.WithCompression(mux), gohm.Config{Timeout: time.Second}),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("[ERROR] ", err)
	}
}
