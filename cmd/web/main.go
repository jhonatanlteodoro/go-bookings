package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jhonatanlteodoro/go-bookings/pkg/config"
	"github.com/jhonatanlteodoro/go-bookings/pkg/handlers"
	"github.com/jhonatanlteodoro/go-bookings/pkg/renders"
)

const portNamber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

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
	}
	app.TemplateCache = tc
	app.UseCache = false

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
