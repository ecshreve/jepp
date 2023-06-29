package models

type Clue struct {
	ID         int64     `gorm:"primary key;autoIncrement" json:"id" example:"804002032"`
	GameID     int64     `json:"gameID" example:"8040"`
	Game       *Game     `gorm:"foreignKey:GameID" json:"game"`
	CategoryID int64     `json:"categoryID" example:"3462"`
	Category   *Category `gorm:"foreignKey:CategoryID" json:"category"`
	Question   string    `json:"question" example:"This is the question."`
	Answer     string    `json:"answer" example:"This is the answer."`
}
