package generic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func Load[T any](
	ctx context.Context,
	db *gorm.DB,
	field string,
	value any,
	out *T,
	preloads ...string,
) (bool, error) {
	q := db.WithContext(ctx).
		Session(&gorm.Session{Logger: db.Logger.LogMode(logger.Silent)})

	for _, preload := range preloads {
		q = q.Preload(preload)
	}

	err := q.Where(field+" = ?", value).First(out).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return false, nil
	case err != nil:
		return false, err
	default:
		return true, nil
	}
}

func Save[T any](ctx context.Context, db *gorm.DB, field string, model *T) error {
	return db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: field}},
			UpdateAll: true,
		}).
		Create(model).Error
}
