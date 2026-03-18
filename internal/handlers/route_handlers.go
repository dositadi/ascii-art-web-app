package handlers

import (
	"context"
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

type AsciiServices interface {
	/* // App pages
	GetAboutUs() ([]m.AboutUs, *m.Error)
	RenderHomePage(w http.ResponseWriter) *m.Error
	RenderLearnMorePage(w http.ResponseWriter) *m.Error

	// Auth pages
	RenderLoginPage(w http.ResponseWriter) *m.Error
	RenderSignupPage(w http.ResponseWriter) *m.Error
	*/

	// Auth Handlers
	RegisterUser(ctx context.Context, name, email, password string) *m.Error
	LoginUser(ctx context.Context, email, password string) (m.ActiveUser, *m.Error)
	CheckDBHealth() *m.Error

	// App Service
	TransformText(request m.Ascii) (string, *m.Error)
}

type Handler struct {
	Service AsciiServices
}

func CreateNewService(service AsciiServices) *Handler {
	return &Handler{
		Service: service,
	}
}

func (s *Handler) AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About connection set!"))
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
	if err := s.Service.CheckDBHealth(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Connection failed!."))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Connection to Db is successful."))
}
