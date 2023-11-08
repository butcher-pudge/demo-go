package controller

import (
	"github.com/gin-gonic/gin"
	"go-learning-demo/auth/constants"
	"go-learning-demo/bookmark/service"
	"go-learning-demo/models"
	"net/http"
)

type BookmarkController struct {
	bookmarkService service.IBookmarkService
}

func NewBookmarkController(bookmarkService service.IBookmarkService) *BookmarkController {
	return &BookmarkController{
		bookmarkService: bookmarkService,
	}
}

type bookmarkCreateRequest struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

func (controller *BookmarkController) CreateBookmark(context *gin.Context) {
	bookmark := new(bookmarkCreateRequest)

	if err := context.BindJSON(bookmark); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := context.MustGet(constants.CtxUserKey).(*models.User)

	if err := controller.bookmarkService.CreateBookmark(context.Request.Context(), user.ID, bookmark.URL, bookmark.Title); err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.Status(http.StatusOK)
}

type bookmarkResponse struct {
	Id    string `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
}

func (controller *BookmarkController) GetBookmarkById(context *gin.Context) {
	bookmarkId := context.Param("id")

	user := context.MustGet(constants.CtxUserKey).(*models.User)

	bookmark, err := controller.bookmarkService.GetBookmarkById(context.Request.Context(), user.ID, bookmarkId)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, toBookmarkResponse(bookmark))
}

func toBookmarkResponse(bookmark *models.Bookmark) *bookmarkResponse {
	return &bookmarkResponse{
		Id:    bookmark.ID,
		URL:   bookmark.URL,
		Title: bookmark.Title,
	}
}
