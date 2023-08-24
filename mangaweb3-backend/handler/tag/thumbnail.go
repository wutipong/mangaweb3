package tag

import (
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/wutipong/mangaweb3-backend/handler"
	"github.com/wutipong/mangaweb3-backend/log"
	"github.com/wutipong/mangaweb3-backend/tag"
	"go.uber.org/zap"
)

func ThumbnailHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tagStr := handler.ParseParam(params, "tag")
	tagStr = filepath.FromSlash(tagStr)

	log.Get().Info("Tag thumbnail image", zap.String("tag", tagStr))

	m, err := tag.Read(r.Context(), tagStr)
	if err != nil {
		handler.WriteResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(m.Thumbnail)
	w.Header().Set("Content-Type", "image/jpeg")
}