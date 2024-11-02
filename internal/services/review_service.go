package services

import "tp/internal/models"

type ReviewService interface {
    CreateReview(review models.Review) (models.Review, error)
    GetReviewsByBookID(bookID int) ([]models.Review, error)
}

type reviewService struct {
    reviewRepo models.ReviewRepository
}

func NewReviewService(reviewRepo models.ReviewRepository) ReviewService {
    return &reviewService{reviewRepo: reviewRepo}
}

func (s *reviewService) CreateReview(review models.Review) (models.Review, error) {
    return s.reviewRepo.Create(review)
}

func (s *reviewService) GetReviewsByBookID(bookID int) ([]models.Review, error) {
    return s.reviewRepo.FindByBookID(bookID)
}

