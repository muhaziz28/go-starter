package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;not null" json:"id"`
	Token     string    `gorm:"type:text;not null" json:"token"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Type      string    `gorm:"type:varchar(50);not null" json:"type"`
	Expires   time.Time `gorm:"not null" json:"expires"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	User      *User     `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

func (token *Token) BeforeCreate(_ *gorm.DB) error {
	if token.ID == uuid.Nil {
		token.ID = uuid.New()
	}
	return nil
}
