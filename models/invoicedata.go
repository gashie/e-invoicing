package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
)

type InvoiceData struct {
	Id               uint           `json:"id" gorm:"primaryKey"`
	Title            string         `json:"title"`
	InvoiceReference string         `json:"invoiceReference" gorm:"unique"`
	InvoiceAppId     uuid.UUID      `gorm:"type:char(36);"`
	InvoiceApps      InvoiceApps    `json:"-" gorm:"foreignKey:InvoiceAppId"`
	ReqDetails       datatypes.JSON `json:"reqDetails" gorm:"type:longtext"`
	Base
}
