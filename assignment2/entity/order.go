package entity

import "time"

type Order struct {
	OrderID      int64     `gorm:"primaryKey:auto_increment" json:"-"`
	CustomerName string    `gorm:"type:varchar(100)" json:"-"`
	OrderAt      time.Time `gorm:"type:date" json:"-"`
	Items        []Items   `json:"items"`
}
