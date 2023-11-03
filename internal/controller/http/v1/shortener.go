package v1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/realPointer/url-shortener/internal/service"
	"github.com/realPointer/url-shortener/pkg/logger"
)

type shortenerRoutes struct {
	shortenerService service.Shortener
	l                logger.Interface
}

func NewShortenerRouter(shortenerService service.Shortener, l logger.Interface) http.Handler {
	s := shortenerRoutes{
		shortenerService: shortenerService,
		l:                l,
	}
	r := chi.NewRouter()

	r.Post("/", s.createShortURL)
	r.Get("/{shortURL}", s.getOriginalURL)

	return r
}

// @Summary Create short URL
// @Description Get short URL from original URL
// @Tags Shortener
// @Accept json
// @Produce json
// @Param url body string true "Original URL"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /shortener [post]
func (s *shortenerRoutes) createShortURL(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	shortURL, err := s.shortenerService.GetShortURL(r.Context(), reqBody.URL)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"shortURL": shortURL})
}

// @Summary Get original URL
// @Description Get original URL from short URL
// @Tags Shortener
// @Param shortURL path string true "Short URL"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /shortener/{shortURL} [get]
func (s *shortenerRoutes) getOriginalURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")

	originalURL, err := s.shortenerService.GetOriginalURL(r.Context(), shortURL)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"originalURL": originalURL})
}
