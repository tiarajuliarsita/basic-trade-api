package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint       `json:"id" form:"id" gorm:"PrimaryKey"`
	UUID      string     `json:"uuid" form:"uuid" gorm:"not null"`
	Name      string     `json:"name" form:"name" gorm:"not null" valid:"required~your product name is required"`
	ImageURL  string     `json:"image_url" form:"image_url" gorm:"not null" valid:"required~your image_url is required"`
	AdminID   uint       `json:"admin_id" form:"admin_id"`
	// Admin     Admin      `json:"admin" form:"admin" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Variants  []Variant  `json:"variant" form:"variant" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	p.UUID = newUuid.String()

	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}
