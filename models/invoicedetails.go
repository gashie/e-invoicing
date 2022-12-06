package models

type InvoiceDetails struct {
	Id                 uint             `json:"id" gorm:"primaryKey"`
	CategoryId         uint             `json:"CategoryId"`
	Category           Category         `json:"-" gorm:"foreignKey:CategoryId"`
	InvoiceDataId      uint             `json:"InvoiceDataId"`
	InvoiceData        InvoiceData      `json:"-" gorm:"foreignKey:InvoiceDataId"`
	InvoiceAppFieldsId uint             `json:"InvoiceAppFieldsId"`
	InvoiceAppFields   InvoiceAppFields `json:"-" gorm:"foreignKey:InvoiceAppFieldsId"`
	FieldValue         string           `json:"fieldVAlue" gorm:"type:varchar(500)"`
	Base
}
