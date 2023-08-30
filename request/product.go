package request

import "mime/multipart"

type Product struct {
	Name string                `json:"name" form:"name" gorm:"not null" valid:"required~your product name is required"`
	File *multipart.FileHeader `form:"file"`
}
