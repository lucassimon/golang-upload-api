package chii

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	dbinfra "github.com/lucassimon/golang-upload-api/internal/adapters/db"
	"github.com/lucassimon/golang-upload-api/internal/adapters/http/chii/handlers"
)

type Webserver struct {
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (web *Webserver) Serve() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	var db, _ = sql.Open("sqlite3", "db.sqlite")
	mediaDB := dbinfra.NewMediaDB(db)

	index_handler := handlers.NewIndexHandler()
	router.Get("/", index_handler.GetIndex)

	upload_v1_handler := handlers.NewUploadV1Handler(mediaDB)
	router.Route("/v1", func(r chi.Router) {
		r.Post("/upload", upload_v1_handler.Upload)
	})

	router.Route("/v2", func(r chi.Router) {
		// r.Post("/upload", upload_v2_handler.Upload)
	})

	err := http.ListenAndServe(":8080", router)
	log.Printf("Listening Port %s...\n", ":8080")
	if err != nil {
		log.Fatal(err)
	}
}
