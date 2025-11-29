package infrastructure

import (
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(shortURL *ShortURLModel) error {
	return r.db.Create(shortURL).Error
}

func (r *Repository) FindByShortCode(shortCode string) (*ShortURLModel, error) {
	var shortURL ShortURLModel
	err := r.db.Where("short_code = ?", shortCode).First(&shortURL).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &shortURL, nil
}