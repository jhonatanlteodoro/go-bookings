package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jhonatanlteodoro/go-bookings/internal/config"
	"github.com/jhonatanlteodoro/go-bookings/internal/handlers"
	"github.com/jhonatanlteodoro/go-bookings/internal/models"
	"github.com/jhonatanlteodoro/go-bookings/internal/renders"
)

const portNamber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	repo := handlers.NewRepo(&app)
	handlers.NewHanlders(repo)
	renders.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNamber,
		Handler: routes(&app),
	}

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNamber))
	// http.ListenAndServe(portNamber, nil)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Fail attempting start server.")
	}
}

func run() error {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in prod
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false

	return nil
}
