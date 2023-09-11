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
	AdminID   uint       
	Admin     *Admin  
	Variants  []Variant 
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {

	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}
	newUuid := uuid.New()
	p.UUID = newUuid.String()

	return nil
}
