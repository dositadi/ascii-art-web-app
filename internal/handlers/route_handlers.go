package handlers

import (
	"context"
	"database/sql"
	"log"
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
	RegisterUser(ctx context.Context, user m.User) *m.Error
	LoginUser(email, password string) *m.Error

	// App Service
	TransformText(request m.Request) (string, *m.Error)
}

type Handler struct {
	Service AsciiServices
	DB      *sql.DB
}

func CreateNewService(service AsciiServices) *Handler {
	return &Handler{
		Service: service,
	}
}

func (s *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login connection set!"))
}

func (s *Handler) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home connection set!"))
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
	if err := s.DB.Ping(); err != nil {
		log.Fatal(err)
	}
	w.Write([]byte("Connection to Db is successful."))
}
