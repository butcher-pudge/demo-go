package mongo

import (
	"context"
	"go-learning-demo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Bookmark struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserId string             `bson:"userId"`
	URL    string             `bson:"url"`
	Title  string             `bson:"title"`
}

type BookmarkRepository struct {
	db *mongo.Collection
}

func NewBookmarkRepository(db *mongo.Database, collection string) *BookmarkRepository {
	return &BookmarkRepository{
		db: db.Collection(collection),
	}
}

func (repository *BookmarkRepository) CreateBookmark(ctx context.Context, bookmark *models.Bookmark) error {
	model := toMongoUser(bookmark)
	res, err := repository.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	bookmark.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (repository *BookmarkRepository) GetBookmarkById(ctx context.Context, userId, bookmarkId string) (*models.Bookmark, error) {
	id, _ := primitive.ObjectIDFromHex(bookmarkId)

	var result *Bookmark
	err := repository.db.FindOne(ctx, bson.M{
		"userId": userId,
		"_id":    id,
	}).Decode(&result)

	return toModel(result), err
}

func toMongoUser(bookmark *models.Bookmark) *Bookmark {
	return &Bookmark{
		UserId: bookmark.UserID,
		URL:    bookmark.URL,
		Title:  bookmark.Title,
	}
}

func toModel(bookmark *Bookmark) *models.Bookmark {
	return &models.Bookmark{
		ID:     bookmark.ID.Hex(),
		UserID: bookmark.UserId,
		URL:    bookmark.URL,
		Title:  bookmark.Title,
	}
}
