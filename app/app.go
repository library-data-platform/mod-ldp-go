package app

import (
	"log"
	"net/http"
	"time"

	"github.com/folio-org/mod-ldp/app/config"
	"github.com/folio-org/mod-ldp/app/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     int // *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	// dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
	// 	config.DB.Username,
	// 	config.DB.Password,
	// 	config.DB.Name,
	// 	config.DB.Charset)

	// db, err := gorm.Open(config.DB.Dialect, dbURI)
	// if err != nil {
	// 	log.Fatal("Could not connect database")
	// }

	// a.DB = model.DBMigrate(db)

	a.DB = 0 // temp value until the database connect code above is done
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/ldp/rt/journal_access_per_time", a.GetJournalAccessPerTime)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage data
func (a *App) GetJournalAccessPerTime(w http.ResponseWriter, r *http.Request) {
	handlers.GetJournalAccessPerTime(a.DB, w, r)
}

// Run the app on its router
func (a *App) Run(host string) {

	srv := &http.Server{
		Handler:      a.Router,
		Addr:         host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
