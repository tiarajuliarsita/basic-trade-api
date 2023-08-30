package models

import (
	"final-project/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint      `json:"id" form:"id" gorm:"PrimaryKey"`
	UUID      string    `json:"uuid" gorm:"not null"`
	Name      string    `json:"name" form:"name" gorm:"not null" valid:"required~your name is required"`
	Email     string    `json:"email" form:"email" gorm:"not null;unique" valid:"required~your email is required, email~invalid email format"`
	Password  string    `json:"password" form:"password" gorm:"not null" valid:"required~your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Products  []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	a.UUID = newUuid.String()

	a.Password, _ = helpers.HassPass(a.Password)
	_, err = govalidator.ValidateStruct(a)
	if err != nil {
		return err
	}

	return nil

}
