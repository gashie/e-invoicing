package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type InvoiceApps struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	AppName     string    `json:"appname"`
	Description string    `json:"description"`
	Base
}

func (u *InvoiceApps) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	return
}
