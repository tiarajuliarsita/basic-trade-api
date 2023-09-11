package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Variant struct {
	ID          uint   `json:"id" form:"id" gorm:"PrimaryKey"`
	UUID        string `json:"uuid" form:"uuid" gorm:"not null"`
	VariantName string `json:"variant_name" form:"variant_name" gorm:"not null" valid:"required~variant~name is required"`
	Quantity    int    `json:"quantity" form:"quantity" gorm:"not null" valid:"required~entity is required"`
	ProductID   uint   `json:"product_id" form:"product_id"`
	Product     *Product
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {

	_, err = govalidator.ValidateStruct(v)
	if err != nil {
		return err
	}
	newUuid := uuid.New()
	v.UUID = newUuid.String()
	return nil
}
