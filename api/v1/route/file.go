package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zetsux/gin-gorm-template-clean/api/v1/controller"
)

func FileRoutes(route *gin.Engine, fileController controller.FileController) {
	routes := route.Group("/api/v1/files")
	{
		routes.GET("/:dir/:file_id", fileController.GetFile)
	}
}
