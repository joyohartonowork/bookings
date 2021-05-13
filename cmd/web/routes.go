package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joyohartonowork/bookings/pkg/config"
	"github.com/joyohartonowork/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("../../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	//println(fmt.Sprintln(fileServer))

	//pages, _ := filepath.Glob("../../static") //tmpl
	//println(fmt.Sprintln(pages))

	// test print current directory files
	// files, err := ioutil.ReadDir("./")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, f := range files {
	// 	fmt.Println(f.Name())
	// }

	return mux
}
