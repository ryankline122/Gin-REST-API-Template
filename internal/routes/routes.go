package routes

import (
    "example/web-service-gin/internal/handlers"
    "github.com/gin-gonic/gin"
)

func SetupRouter(albumHandler *handlers.AlbumHandler) *gin.Engine {
    router := gin.Default()

    albumGroup := router.Group("albums/")
    {
        albumGroup.GET("/", albumHandler.GetAlbums)
        albumGroup.GET("/:id", albumHandler.GetAlbumByID)
        albumGroup.POST("/", albumHandler.AddAlbum)
        albumGroup.POST("/update/:id/", albumHandler.UpdateAlbumPrice)
        albumGroup.DELETE("/:id", albumHandler.DeleteAlbumByID)
    }

    return router
}
