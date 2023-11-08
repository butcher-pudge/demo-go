package repository

import (
	"context"
	"go-learning-demo/models"
)

type IBookmarkRepository interface {
	CreateBookmark(ctx context.Context, bookmark *models.Bookmark) error
	GetBookmarkById(ctx context.Context, userId, bookmarkId string) (*models.Bookmark, error)
}
