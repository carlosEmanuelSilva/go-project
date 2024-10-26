package models

import "gorm.io/gorm"

type User struct {
  ID        uint    `gorm: "primaryKey;autoIncrement" json:"id"`
  Name      string  `gorm: "size:100;not null" json:"name"`
  Email     string  `gorm: "size:100;unique;not null" json:"email"`
  Password  string  `gorm: "not null" json:"-"`
}
