package scheduler

import (
	"context"

	"github.com/wutipong/mangaweb3-backend/log"
	"github.com/wutipong/mangaweb3-backend/meta"
)

func ScanLibrary() error {
	allMeta, err := meta.ReadAll(context.Background())
	if err != nil {
		return err
	}

	files, err := meta.ListDir("")
	if err != nil {
		return err
	}

	for _, file := range files {
		found := false
		for _, m := range allMeta {
			if m.Name == file {
				found = true
				break
			}
		}
		if found {
			continue
		}

		log.Get().Sugar().Infof("Creating metadata for %s", file)

		item, err := meta.NewItem(file)
		if err != nil {
			log.Get().Sugar().Errorf("Failed to create meta data : %v", err)
		}

		err = meta.Write(context.Background(), item)
		if err != nil {
			log.Get().Sugar().Errorf("Failed to write meta data : %v", err)
		}
	}

	for _, m := range allMeta {
		found := false
		for _, file := range files {
			if m.Name == file {
				found = true
				break
			}
		}
		if found {
			continue
		}

		log.Get().Sugar().Infof("Deleting metadata for %s", m.Name)
		if err := meta.Delete(context.Background(), m); err != nil {
			log.Get().Sugar().Infof("Failed to delete meta for %s", m.Name)
		}

	}

	return nil
}

func ScheduleScanLibrary() {
	scheduler.Every(1).Millisecond().LimitRunsTo(1).Do(func() {
		log.Get().Sugar().Infof("Scanning Library.")
		ScanLibrary()
	})
}