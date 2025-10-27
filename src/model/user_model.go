// ========================================
// FILE: src/model/user.go
// ========================================
package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;not null" json:"id"`
	Name          string    `gorm:"type:varchar(255);not null" json:"name"`
	Email         string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password      string    `gorm:"type:varchar(255);not null" json:"-"`
	VerifiedEmail bool      `gorm:"type:boolean;default:false;not null" json:"verified_email"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Token         []Token   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

func (user *User) BeforeCreate(_ *gorm.DB) error {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	return nil
}
