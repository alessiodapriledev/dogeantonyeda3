package models

type Alert struct {
	ID      uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Lat     float64 `json:"lat" gorm:"not null"`
	Lng     float64 `json:"lng" gorm:"not null"`
	Message string  `json:"message" gorm:"type:text;not null"`
}
