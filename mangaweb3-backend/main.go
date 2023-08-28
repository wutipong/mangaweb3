package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "github.com/wutipong/mangaweb3-backend/docs"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/handler/browse"
	handlertag "github.com/wutipong/mangaweb3-backend/handler/tag"
	"github.com/wutipong/mangaweb3-backend/handler/view"
	"github.com/wutipong/mangaweb3-backend/meta"
	"github.com/wutipong/mangaweb3-backend/scheduler"
	"github.com/wutipong/mangaweb3-backend/tag"
)

var versionString string = "development"

//go:generate go run -mod=mod github.com/swaggo/swag/cmd/swag@latest init

// @title           Mangaweb3 API
// @version         3.0
// @description     API Server for Mangaweb

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	if err := godotenv.Overload(); err == nil {
		log.Info().Msg("Use .env file.")
	}

	address := ":8972"
	if v, b := os.LookupEnv("MANGAWEB_ADDRESS"); b {
		address = v
	}

	dataPath := "./data"
	if v, b := os.LookupEnv("MANGAWEB_DATA_PATH"); b {
		dataPath = v
	}
	connectionStr := "postgres://postgres:password@localhost:5432/manga"
	if v, b := os.LookupEnv("MANGAWEB_DB"); b {
		connectionStr = v
	}

	if _, b := os.LookupEnv("MANGAWEB_DEVELOPMENT"); b {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).
			Level(zerolog.DebugLevel)
	}

	meta.BaseDirectory = dataPath

	log.Info().
		Str("version", versionString).
		Str("data_path", dataPath).
		Str("address", address).
		Msg("Server started.")

	var client *ent.Client = nil
	if db, err := sql.Open("pgx", connectionStr); err != nil {
		log.Error().AnErr("error", err).Msg("Connect to Postgres fails")
		return
	} else {
		drv := entsql.OpenDB(dialect.Postgres, db)
		defer db.Close()

		client = ent.NewClient(ent.Driver(drv))
		defer client.Close()
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Error().AnErr("error", err).Msg("failed creating schema resources.")
		return
	}

	meta.Init(client)
	tag.Init(client)

	scheduler.Init(scheduler.Options{})
	scheduler.Start()

	router := httprouter.New()
	RegisterHandler(router)

	log.Info().Msg("Server starts.")

	handler := cors.Default().Handler(router)
	if err := http.ListenAndServe(address, handler); err != nil {
		log.Error().AnErr("error", err).Msg("Starting server fails")
		return
	}

	log.Info().Msg("shutting down the server")
	scheduler.Stop()
}

func RegisterHandler(router *httprouter.Router) {
	handler.Init(handler.Options{
		VersionString: versionString,

		PathView:        view.PathView,
		PathGetImage:    view.PathGetImage,
		PathUpdateCover: view.PathUpdateCover,
		PathFavorite:    view.PathFavorite,
		PathDownload:    view.PathDownload,

		PathThumbnail:     browse.PathThumbnail,
		PathRescanLibrary: browse.PathRescanLibrary,
		PathBrowse:        browse.PathBrowse,

		PathTagFavorite:  handlertag.PathSetFavorite,
		PathTagList:      handlertag.PathList,
		PathTagThumbnail: handlertag.PathThumbnail,
	})
	// Routes
	router.GET(browse.PathRescanLibrary, browse.RescanLibraryHandler)
	router.GET(browse.PathThumbnail, browse.GetThumbnailHandler)
	router.POST(browse.PathBrowse, browse.Handler)
	router.GET(browse.PathRecreateThumbnails, browse.RecreateThumbnailHandler)

	router.GET(view.PathDownload, view.Download)
	router.GET(view.PathGetImage, view.GetImage)
	router.POST(view.PathFavorite, view.SetFavoriteHandler)
	router.POST(view.PathUpdateCover, view.UpdateCover)
	router.POST(view.PathView, view.Handler)

	router.GET(handlertag.PathThumbnail, handlertag.ThumbnailHandler)
	router.POST(handlertag.PathList, handlertag.ListHandler)
	router.POST(handlertag.PathSetFavorite, handlertag.SetFavoriteHandler)

	router.GET("/doc/:any", swaggerHandler)
}

func swaggerHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	httpSwagger.WrapHandler(res, req)
}
