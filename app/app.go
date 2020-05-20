package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/folio-org/mod-ldp/app/config"
	"github.com/folio-org/mod-ldp/app/handlers"
	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func SetupCloseHandler(db *gorm.DB) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		db.Close()
		os.Exit(0)
	}()
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("host=%s dbname=ldp_folio_release port=%s user=%s password=%s dbname=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	} else {
		log.Println("Connected to db")
	}
	SetupCloseHandler(db)

	// type Result struct {
	// 	ID   string
	// 	Type string
	// }
	// db.AutoMigrate(&Result{})

	// var result Result
	// db.Raw("SELECT id,type FROM public.user_users LIMIT 1").Scan(&result)
	// fmt.Println(result)
	// a.DB = model.DBMigrate(db)

	a.DB = db // temp value until the database connect code above is done
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/ldp/rt/journal_access_per_time", a.GetJournalAccessPerTime)
	a.Get("/ldp/db/status", a.GetDbStatus)
	a.Get("/ldp/db/log", a.GetLogTable)
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

func (a *App) GetDbStatus(w http.ResponseWriter, r *http.Request) {
	handlers.GetDbStatus(a.DB, w, r)
}
func (a *App) GetLogTable(w http.ResponseWriter, r *http.Request) {
	handlers.GetLogTable(a.DB, w, r)
}

// Run the app on its router
func (a *App) Run(host string) {

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := ghandlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := ghandlers.AllowedOrigins([]string{"*"}) // TODO: use os.Getenv("ORIGIN_ALLOWED")
	methodsOk := ghandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	srv := &http.Server{
		Handler:      ghandlers.CORS(headersOk, originsOk, methodsOk)(a.Router),
		Addr:         host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
