package models

import (
	uuid "github.com/satori/go.uuid"
)

type InvoiceAppFields struct {
	Id uint `json:"id" gorm:"primaryKey"`

	Name         string      `json:"name"`
	DataType     string      `json:"dataType"`
	FieldLabel   string      `json:"fieldLabel"`
	CategoryId   uint        `json:"CategoryId"`
	Category     Category    `json:"-" gorm:"foreignKey:CategoryId"`
	InvoiceAppId uuid.UUID   `gorm:"type:char(36);"`
	InvoiceApps  InvoiceApps `json:"-" gorm:"foreignKey:InvoiceAppId"`
	Description  string      `json:"description"`
	Hide         bool        `json:"hide" gorm:"type:bool;default:false"`
	IsContact    bool        `json:"isContact" gorm:"type:bool;default:false"`
	IsEmail      bool        `json:"isEmail" gorm:"type:bool;default:false"`
	IsAmount     bool        `json:"isAmount" gorm:"type:bool;default:false"`
	IsTax        bool        `json:"isTax" gorm:"type:bool;default:false"`
	Base
}
