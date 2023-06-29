package models

type Category struct {
	ID   int64  `gorm:"primary key;autoIncrement" json:"id" example:"3462"`
	Name string `json:"name" example:"HISTORY"`
}
