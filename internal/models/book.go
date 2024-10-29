package models

import "gorm.io/gorm"

type Book struct{
  gorm.Model
  ID            uint    `gorm:"primaryKey;autoIncrement" json:"id"`
  Title         string  `gorm:"size:200;not null" json:"title"`
  Author        string  `gorm:"size:100;not null" json:"author"`
  Genre         string  `gorm:"size:50;not null" json:"genre"` 
  PublishYear   int     `gorm:"not null" json"published_year"`
  AvarageRating float64 `gorm:"deafault:0" json"avarage_rating"`
}
