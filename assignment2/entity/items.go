package entity

type Items struct {
	ItemID      int64  `gorm:"primaryKey:auto_increment" json:"-"`
	ItemCode    string `gorm:"type:varchar(100)" json:"-"`
	Description string `gorm:"type:text" json:"-"`
	Quantity    int    `gorm:"type:integer" json:"-"`
	OrderID     int    `gorm:"type:integer" json:"-"`
	Order       Order  `json:"-"`
}
