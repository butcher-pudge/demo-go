package bookmark

import (
	"github.com/gin-gonic/gin"
	"go-learning-demo/bookmark/controller"
)

func RegisterHttpEndpoints(router *gin.RouterGroup, controller *controller.BookmarkController) {

	authEndpoints := router.Group("/bookmark")
	{
		authEndpoints.POST("", controller.CreateBookmark)
		authEndpoints.GET("/:id", controller.GetBookmarkById)
	}
}
