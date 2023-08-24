package scheduler

import (
	"context"

	"github.com/wutipong/mangaweb3-backend/log"
	"github.com/wutipong/mangaweb3-backend/meta"
)

func RebuildThumbnail() error {
	allMeta, err := meta.ReadAll(context.Background())
	if err != nil {
		return err
	}

	for _, m := range allMeta {
		e := m.GenerateThumbnail(0)
		log.Get().Sugar().Infof("Generating new thumbnail for %s", m.Name)
		if e != nil {
			log.Get().Sugar().Errorf("Failed to generate thumbnail for %s", m.Name)
			continue
		}

		meta.Write(context.Background(), m)
	}

	return nil
}

func ScheduleRebuildThumbnail() {
	scheduler.Every(1).Millisecond().LimitRunsTo(1).Do(func() {
		log.Get().Sugar().Info("Force updating thumbnail")
		RebuildThumbnail()
	})
}