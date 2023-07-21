package models

type Movie struct {
	ID            uint   `gorm:"primary_key"`
	Name          string `gorm:"type:varchar(100)"`
	ReleaseYear   int
	Director      string `gorm:"type:varchar(100)"`
	Actors        string `gorm:"type:varchar(500)"` // This is a comma-separated list of actors
	AverageRating float64
	Duration      int
	Genre         string `gorm:"type:varchar(100)"`
}
