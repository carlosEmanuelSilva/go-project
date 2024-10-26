package models

import "gorm.io/gorm"

type Review struct {
    ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID  uint   `gorm:"not null" json:"user_id"`
    BookID  uint   `gorm:"not null" json:"book_id"`
    Rating  int    `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating"`
    Comment string `gorm:"type:text" json:"comment"`
    
    User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
    Book Book `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE" json:"-"`
}

