package urls

import (
	"net/http"
	"shrinklink/service/urls"

	"github.com/gorilla/mux"
)

type UrlsHandler struct {
	service urls.IUrlService
}

func NewUrlsHandler(service urls.IUrlService) *UrlsHandler {
	handler := &UrlsHandler{service: service}
	return handler
}

func (h *UrlsHandler) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/urls", h.GetAllUrls).Methods(http.MethodGet)
	r.HandleFunc("/urls", h.AddUrl).Methods(http.MethodPost)
	r.HandleFunc("/url/{short_url}", h.GetUrl).Methods(http.MethodGet)
}
