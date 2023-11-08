package service

import (
	"context"
	"go-learning-demo/models"
)

type IBookmarkService interface {
	CreateBookmark(ctx context.Context, userId, url, title string) error
	GetBookmarkById(ctx context.Context, userId, bookmarkId string) (*models.Bookmark, error)
}
