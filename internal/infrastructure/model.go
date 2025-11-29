package infrastructure

import (
	"time"

	"github.com/google/uuid"
)

type ShortURL struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	LongUrl string  `gorm:"type:text;not null"`
	ShortCode   string  `gorm:"type:varchar(20);uniqueIndex;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (ShortURL) TableName() string {
	return "short_urls"
}
