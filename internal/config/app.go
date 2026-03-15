package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	h_ "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/handlers"
	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	DB             *sql.DB
	Router         *http.ServeMux
	ServiceHandler h_.Handler
}

func (a *App) InitializeFileServers() {
	pagesCSSFileServer := http.FileServer(http.Dir(h.STYLES_PATH))
	pattern := "GET " + h.STYLES_PATH_PATTERN
	a.Router.Handle(pattern, http.StripPrefix(h.STYLES_PATH_PATTERN, pagesCSSFileServer))
}

func (a *App) InitializeRoutes() {
	a.Router = http.NewServeMux()

	// Welcome Page
	a.Router.HandleFunc("GET /", a.ServiceHandler.WelcomePageHandler)

	// Auth route
	a.Router.HandleFunc("GET /auth/login", a.ServiceHandler.LoginPageHandler)
	a.Router.HandleFunc("POST /auth/login", a.ServiceHandler.LoginPageHandler)
	a.Router.HandleFunc("POST /auth/register", a.ServiceHandler.RegisterHandler)
	a.Router.HandleFunc("GET /auth/register", a.ServiceHandler.RegisterPageHandler)

	// Pages route
	a.Router.HandleFunc("GET /ascii-art/learn-more", a.ServiceHandler.LearnMorePageHandler)
	a.Router.HandleFunc("GET /ascii-art/about-us", a.ServiceHandler.AboutPageHandler)

	a.Router.HandleFunc("GET /health", a.ServiceHandler.HealthCheckHandler)
}

func (a *App) InitializeDatabase() *m.Error {
	c := Config{}
	config := c.GetDBConfig()

	var err error

	mysqlConfig := mysql.NewConfig()
	mysqlConfig.Addr = config.Addr
	mysqlConfig.DBName = config.DBName
	mysqlConfig.Passwd = config.Password
	mysqlConfig.User = config.User
	mysqlConfig.Net = "tcp"

	fmt.Println(mysqlConfig.Addr, mysqlConfig.DBName, mysqlConfig.Passwd, mysqlConfig.User)

	a.DB, err = sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: h.SERVER_ERR_DETAIL,
			Code:    h.SERVER_ERR_CODE,
		}
	}

	a.ServiceHandler.DB = a.DB

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
		Addr:              ":8081",
		Handler:           a.Router,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
