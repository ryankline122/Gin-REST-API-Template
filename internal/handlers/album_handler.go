package handlers

import (
    "example/web-service-gin/internal/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

type AlbumHandler struct {
    AlbumService *services.AlbumService
}

func (handler *AlbumHandler) GetAlbums(context *gin.Context) {
    albums := handler.AlbumService.GetAllAlbums()
    context.JSON(http.StatusOK, albums)
}
