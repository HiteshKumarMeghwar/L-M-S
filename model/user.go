package model

import "gorm.io/gorm"

type User struct {
	Id          int    `gorm:"type:int;primary_key" json:"id"`
	Name        string `gorm:"type:varchar(50);not null" json:"name"`
	Author      string `gorm:"type:varchar(50);not null" json:"author"`
	Description string `gorm:"type:varchar(50);not null" json:"description"`
	*gorm.Model
}
