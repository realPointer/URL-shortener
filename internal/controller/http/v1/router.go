package v1

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/realPointer/url-shortener/internal/service"
	"github.com/realPointer/url-shortener/pkg/logger"

	_ "github.com/realPointer/url-shortener/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewRouter(handler chi.Router, l logger.Interface, services *service.Services) {
	handler.Use(middleware.Logger)
	handler.Use(middleware.Recoverer)
	handler.Use(middleware.Timeout(60 * time.Second))

	handler.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong!"))
	})

	handler.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	handler.Route("/v1", func(r chi.Router) {
		r.Mount("/shortener", NewShortenerRouter(services.Shortener, l))
	})
}
