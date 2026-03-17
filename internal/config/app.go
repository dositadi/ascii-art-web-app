package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	h_ "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/handlers"
	mid "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/middleware"
	r "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/repository"
	s "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/services"
	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	DB     *sql.DB
	Router *http.ServeMux
}

func (a *App) InitializeFileServers() {
	pagesCSSFileServer := http.FileServer(http.Dir(h.STYLES_PATH))
	pattern := "GET " + h.STYLES_PATH_PATTERN
	a.Router.Handle(pattern, http.StripPrefix(h.STYLES_PATH_PATTERN, pagesCSSFileServer))
}

func (a *App) InitializeRoutes() {
	a.Router = http.NewServeMux()

	dB := r.ConstructNewRepo(a.DB)
	service := s.ConstructNewService(dB)
	handler := h_.CreateNewService(service)

	// Welcome Page
	a.Router.HandleFunc("GET /", handler.WelcomePageHandler)

	// Auth route
	a.Router.HandleFunc("GET /auth/login", handler.LoginPageHandler)
	a.Router.HandleFunc("POST /auth/login", handler.LoginHandler)

	a.Router.HandleFunc("POST /auth/register", handler.RegisterHandler)
	a.Router.HandleFunc("GET /auth/register", handler.RegisterPageHandler)

	// Pages route
	a.Router.Handle("GET /home", mid.AuthenticateUser(http.HandlerFunc(handler.HomePageHandler)))
	a.Router.HandleFunc("GET /ascii-art/about-us", handler.AboutPageHandler)

	a.Router.HandleFunc("GET /health", handler.HealthCheckHandler)
}

func (a *App) InitializeDatabase() *m.Error {
	c := Config{}
	config := c.GetDBConfig()

	var err error

	a.DB, err = sql.Open("mysql", config.DBUrl)
	if err != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: h.SERVER_ERR_DETAIL,
			Code:    h.SERVER_ERR_CODE,
		}
	}

	a.DB.SetMaxOpenConns(10)
	a.DB.SetMaxIdleConns(5)
	a.DB.SetMaxOpenConns(20)
	a.DB.SetConnMaxLifetime(time.Duration(10 * time.Minute))

	a.InitializeRoutes()
	a.InitializeFileServers()
	return nil
}

func (a *App) Run() {
	if err := a.InitializeDatabase(); err != nil {

	}

	server := http.Server{
		Addr:              ":8082",
		Handler:           a.Router,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	fmt.Println("Server is running now!.")

	log.Fatal(server.ListenAndServe())
}
