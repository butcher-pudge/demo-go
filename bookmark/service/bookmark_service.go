package service

import (
	"context"
	"go-learning-demo/bookmark/repository"
	"go-learning-demo/models"
)

type BookmarkService struct {
	bookmarkRepository repository.IBookmarkRepository
}

func NewBookmarkService(bookmarkRepository repository.IBookmarkRepository) *BookmarkService {
	return &BookmarkService{
		bookmarkRepository: bookmarkRepository,
	}
}

func (service *BookmarkService) CreateBookmark(ctx context.Context, userId, url, title string) error {
	bookmark := &models.Bookmark{
		UserID: userId,
		URL:    url,
		Title:  title,
	}

	return service.bookmarkRepository.CreateBookmark(ctx, bookmark)
}

func (service *BookmarkService) GetBookmarkById(ctx context.Context, userId, bookmarkId string) (*models.Bookmark, error) {
	return service.bookmarkRepository.GetBookmarkById(ctx, userId, bookmarkId)
}
