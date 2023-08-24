package view

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/log"
	"github.com/wutipong/mangaweb3-backend/meta"
	"go.uber.org/zap"
)

type setFavoriteResponse struct {
	Favorite bool `json:"favorite"`
}

func SetFavoriteHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	item := handler.ParseParam(params, "item")
	item = filepath.FromSlash(item)

	log.Get().Info("Set Favorite Item", zap.String("item_name", item))

	query := r.URL.Query()

	m, err := meta.Read(r.Context(), item)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	if fav, e := strconv.ParseBool(query.Get("favorite")); e == nil {
		if fav != m.Favorite {
			m.Favorite = fav
			meta.Write(r.Context(), m)
		}
	}

	response := setFavoriteResponse{
		Favorite: m.Favorite,
	}

	handler.WriteResponse(w, response)
}