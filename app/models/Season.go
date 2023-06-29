package models

import "time"

type Season struct {
	ID        int64     `gorm:"primary key;autoIncrement" json:"id" example:"38"`
	Number    int64     `json:"number" example:"38"`
	StartDate time.Time `json:"startDate" example:"2019-01-01"`
	EndDate   time.Time `json:"endDate" example:"2019-01-01"`
}
