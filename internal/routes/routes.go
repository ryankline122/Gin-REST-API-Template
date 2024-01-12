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
        // Add more routes as needed
    }

    return router
}
