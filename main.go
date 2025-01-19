package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"templ-components/handler"
	"templ-components/static"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	dsn  = flag.String("dsn", ":memory:", "sqlite dsn")
	port = flag.Int("port", 4000, "http server port")
)

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(noCacheMiddleware)

	homeHandler := handler.NewHome()

	r.Group(func(r chi.Router) {
		r.Get("/", homeHandler.Index)
	})

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServerFS(static.FS)))

	s := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: r,
	}

	slog.Info(fmt.Sprintf("http://localhost:%d", *port))

	log.Fatal(s.ListenAndServe())
}
