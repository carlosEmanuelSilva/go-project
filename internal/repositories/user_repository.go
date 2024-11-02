package models

import "gorm.io/gorm"

type UserRepository interface {
    Create(user User) (User, error)
    FindByID(id int) (User, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(user User) (User, error) {
    if err := r.db.Create(&user).Error; err != nil {
        return User{}, err
    }
    return user, nil
}

func (r *userRepository) FindByID(id int) (User, error) {
    var user User
    if err := r.db.First(&user, id).Error; err != nil {
        return User{}, err
    }
    return user, nil
}

