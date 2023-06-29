package models

import "time"

// Game represents a single game of Jeopardy.
type Game struct {
	ID       int64     `gorm:"primary key;autoIncrement" json:"id" example:"8040"`
	SeasonID int64     `json:"seasonID" example:"38"`
	Season   *Season   `gorm:"foreignKey:SeasonID" json:"season"`
	Show     int64     `json:"showNum" example:"4532"`
	AirDate  time.Time `json:"airDate" example:"2019-01-01"`
	TapeDate time.Time `json:"tapeDate" example:"2019-01-01"`
}
