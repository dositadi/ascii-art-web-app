package config

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	h_ "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/handlers"
	mid "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/middleware"
	r "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/repository"
	s "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/services"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type App struct {
	DB     *pgxpool.Pool
	Router *http.ServeMux
}

func (a *App) InitializeFileServers() {
	pagesCSSFileServer := http.FileServer(http.Dir(h.STYLES_PATH))
	pattern := "GET " + h.STYLES_PATH_PATTERN
	a.Router.Handle(pattern, http.StripPrefix(h.STYLES_PATH_PATTERN, pagesCSSFileServer))

	webImageFileServer := http.FileServer(http.Dir("web/assets"))
	a.Router.Handle("GET /web/assets/", http.StripPrefix("/web/assets/", webImageFileServer))
}

func (a *App) Initialize() {
	a.InitializeDatabase()
	a.Router = http.NewServeMux()

	dB := r.ConstructNewRepo(a.DB)
	service := s.ConstructNewService(dB)
	handler := h_.CreateNewService(service)

	// Welcome Page
	a.Router.HandleFunc("GET "+h.WELCOME_ROUTE, handler.WelcomePageHandler)

	// Auth route
	a.Router.HandleFunc("GET "+h.LOGIN_ROUTE, handler.LoginPageHandler)
	a.Router.HandleFunc("POST "+h.LOGIN_ROUTE, handler.LoginHandler)

	a.Router.HandleFunc("POST "+h.SIGNUP_ROUTE, handler.RegisterHandler)
	a.Router.HandleFunc("GET "+h.SIGNUP_ROUTE, handler.RegisterPageHandler)
	a.Router.HandleFunc("GET "+h.SESSION_EXPIRED_ROUTE, handler.SessionExpiredHandler)
	a.Router.Handle("GET "+h.LOGOUT_ROUTE, mid.AuthenticateUser(http.HandlerFunc(handler.LogoutHandler)))

	// Pages route
	a.Router.Handle("GET "+h.HOME_ROUTE, mid.AuthenticateUser(http.HandlerFunc(handler.HomePageHandler)))
	a.Router.Handle("GET "+h.ASCII_ROUTE, mid.AuthenticateUser(http.HandlerFunc(handler.AsciiArtPageHandler)))
	a.Router.Handle("GET "+h.HISTORY_ROUTE, mid.AuthenticateUser(http.HandlerFunc(handler.HistoryPageHandler)))

	// Text transform route
	a.Router.HandleFunc("POST "+h.ASCII_ROUTE, handler.TransformTextHandler)
	a.Router.Handle("POST "+h.SAVE_ASCII_ROUTE, mid.AuthenticateUser(http.HandlerFunc(handler.SaveAsciiHandler)))
	a.Router.Handle("DELETE "+h.DELETE_ROUTE, mid.AuthenticateUser(http.HandlerFunc(handler.DeleteAsciiHandler)))
	a.Router.Handle("GET "+h.CLEAR_ALL_ROUTE, mid.AuthenticateUser(http.HandlerFunc(handler.ClearAllHandler)))
	a.Router.Handle("POST "+h.COPY_ASCII_ROUTE, mid.AuthenticateUser(http.HandlerFunc(handler.CopyHandler)))
	a.Router.Handle("GET "+h.DOWNLOAD_ASCII_AS_TXT_ROUTE, http.HandlerFunc(handler.DownloadAsciiTxtHandler))

	a.Router.HandleFunc("GET /health", handler.HealthCheckHandler)
	a.InitializeFileServers()
}

func (a *App) InitializeDatabase() {
	a.DB = a.connectToDB()
	if a.DB == nil {
		log.Fatal("Unable to connect to the database!")
		return
	}

	ctx, cancelPing := context.WithTimeout(context.Background(), time.Duration(2*time.Minute))

	defer cancelPing()

	err := a.DB.Ping(ctx)
	if err != nil {
		log.Fatal("Unable to connect to the database!: ", err)
		return
	}

	fmt.Println("DB connected!")
}

func (a *App) connectToDB() *pgxpool.Pool {
	godotenv.Load()

	url := os.Getenv("POSTGRES_DB_URL")
	var err error
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		// handle error
	}
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	for i := 0; i < 10; i++ {
		pool, err := pgxpool.NewWithConfig(context.Background(), config)

		if err == nil {
			return pool
		}

		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (a *App) Run() {
	a.Initialize()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	server := http.Server{
		Addr:              ":" + port,
		Handler:           a.Router,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	fmt.Println("Server is running now!.")

	log.Fatal(server.ListenAndServe())
}
