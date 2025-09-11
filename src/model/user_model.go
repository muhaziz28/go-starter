package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            uuid.UUID `gorm:"primaryKey;type:char(36);not null" json:"id"`
	Name          string    `gorm:"type:varchar(255);not null" json:"name"`
	Email         string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password      string    `gorm:"type:varchar(255);not null" json:"-"`
	VerifiedEmail bool      `gorm:"type:boolean;default:false;not null" json:"verified_email"`
	CreatedAt     time.Time `gorm:"type:datetime;autoCreateTime" json:"-"`
	UpdatedAt     time.Time `gorm:"type:datetime;autoCreateTime;autoUpdateTime" json:"-"`
	Token         []Token   `gorm:"foreignKey:UserID;references:ID" json:"-"`
}

func (user *User) BeforeCreate(_ *gorm.DB) error {
	user.ID = uuid.New() // Generate UUID before create
	return nil
}
