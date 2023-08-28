package scheduler

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/meta"
)

func RebuildThumbnail() error {
	allMeta, err := meta.ReadAll(context.Background())
	if err != nil {
		return err
	}

	for _, m := range allMeta {
		e := meta.GenerateThumbnail(m, 0)
		log.Info().Str("name", m.Name).Msg("Create new thumbnail")
		if e != nil {
			log.Error().Str("item", m.Name).AnErr("error", e).Msg("Fails to create thumbnail.")
			continue
		}

		meta.Write(context.Background(), m)
	}

	return nil
}

func ScheduleRebuildThumbnail() {
	scheduler.Every(1).Millisecond().LimitRunsTo(1).Do(func() {
		log.Info().Msg("Force updating thumbnail")
		RebuildThumbnail()
	})
}
