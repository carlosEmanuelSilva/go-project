package models

import "gorm.io/gorm"

type ReviewRepository interface {
    Create(review Review) (Review, error)
    FindByBookID(bookID int) ([]Review, error)
}

type reviewRepository struct {
    db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
    return &reviewRepository{db: db}
}

func (r *reviewRepository) Create(review Review) (Review, error) {
    if err := r.db.Create(&review).Error; err != nil {
        return Review{}, err
    }
    return review, nil
}

func (r *reviewRepository) FindByBookID(bookID int) ([]Review, error) {
    var reviews []Review
    if err := r.db.Where("book_id = ?", bookID).Find(&reviews).Error; err != nil {
        return nil, err
    }
    return reviews, nil
}

