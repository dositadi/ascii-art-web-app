package handlers

import (
	"context"
	"net/http"
	"time"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

type AsciiServices interface {
	// Auth Handlers
	RegisterUser(ctx context.Context, name, email, password string) *m.Error
	LoginUser(ctx context.Context, email *string, password string) (m.ActiveUser, *m.Error)
	CheckDBHealth(ctx context.Context) *m.Error

	// Render Page functions
	RenderAsciiArtPage(w http.ResponseWriter, r *http.Request) *m.Error
	RenderWelcomePage(w http.ResponseWriter, r *http.Request) *m.Error
	RenderLoginPage(w http.ResponseWriter, r *http.Request, message *string) *m.Error
	RenderSignupPage(w http.ResponseWriter, r *http.Request, message *string) *m.Error
	RenderHomePage(w http.ResponseWriter, r *http.Request) *m.Error
	RenderHistoryPage(w http.ResponseWriter, r *http.Request, font string) *m.Error
	RenderSessionExpiredPage(w http.ResponseWriter, r *http.Request) *m.Error

	// Render HTML partials function
	//RenderAsciiPlaceholder(w http.ResponseWriter, r *http.Request) *m.Error

	// App Service
	TransformText(w http.ResponseWriter, r *http.Request, text, banner string, start time.Time) *m.Error
	SaveAscii(ctx context.Context, text, banner, user_id string) *m.Error
	DeleteAscii(ctx context.Context, id string) *m.Error
	ClearAllSavedAscii(ctx context.Context, user_id string) *m.Error
	CopyAscii(w http.ResponseWriter, r *http.Request) *m.Error
	DownloadAsTxt(w http.ResponseWriter, r *http.Request, text, font, id string) *m.Error
}

type Handler struct {
	Service AsciiServices
}

func CreateNewService(service AsciiServices) *Handler {
	return &Handler{
		Service: service,
	}
}

func (s *Handler) LearnMorePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Learn More connection set!"))
}

func (s *Handler) ServerErrorPageHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Value", "text/html")
	w.Write([]byte("Internal Server Error Ocurred."))
}

func (s *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.Service.CheckDBHealth(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Connection failed!."))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Connection to Db is successful."))
}

func (h *Handler) GetUserID(r *http.Request) string {
	val := r.Context().Value("user_id")

	if id, ok := val.(string); ok {
		return id
	}
	return ""
}
