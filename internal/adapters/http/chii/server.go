package chii

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	var db, _ = sql.Open("sqlite3", "db.sqlite")
	mediaDB := dbinfra.NewMediaDB(db)

	index_handler := handlers.NewIndexHandler()
	router.Get("/", index_handler.GetIndex)

	upload_v1_handler := handlers.NewUploadV1Handler(mediaDB)
	router.Route("/v1", func(r chi.Router) {
		// https://github.com/go-chi/chi/blob/master/_examples/versions/main.go
		// r.Use(apiVersionCtx("v1"))
		r.Post("/upload", upload_v1_handler.Upload)
	})

	// upload_v2_handler := handlers.NewUploadV2Handler(mediaDB)
	// router.Route("/v2", func(r chi.Router) {
	// 	r.Use(apiVersionCtx("v2"))
	// 	r.Post("/upload", upload_v2_handler.Upload)
	// })

	medias_handler := handlers.NewMediasHandler(mediaDB)
	router.Route("/admin", func(r chi.Router) {
		r.Get("/medias", medias_handler.GetAll)
		r.Route("/medias/{id}", func(r chi.Router) {
			r.Get("/", medias_handler.GetById)
			r.Delete("/", medias_handler.Delete)
			r.Get("/check", medias_handler.Sync)
		})
	})

	// FileServer
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "media"))
	FileServer(router, "/files", filesDir)

	// HTTP server
	err := http.ListenAndServe(":8080", router)
	log.Printf("Listening Port %s...\n", ":8080")
	if err != nil {
		log.Fatal(err)
	}
}

// func apiVersionCtx(version string) func(next http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

// FileServer is serving static files.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
