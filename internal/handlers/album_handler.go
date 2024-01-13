package handlers

import (
	"example/web-service-gin/internal/models"
	"example/web-service-gin/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
    AlbumService *services.AlbumService
}

func (h *AlbumHandler) AddAlbum(c *gin.Context) {
    var newAlbum models.Album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    h.AlbumService.AddAlbum(newAlbum)
    c.JSON(http.StatusCreated, newAlbum)
}

func (h *AlbumHandler) GetAlbums(c *gin.Context) {
    albums, err := h.AlbumService.GetAllAlbums()
    if err != nil {
        c.JSON(http.StatusNotFound, err)
    }
    c.JSON(http.StatusOK, albums)
}

func (h *AlbumHandler) GetAlbumByID(c *gin.Context) {
    id := c.Param("id")
    album, err := h.AlbumService.GetAlbumByID(id)
    
    if err != nil {
        c.JSON(http.StatusNotFound, err)
    }
    
    c.JSON(http.StatusOK, album)
}

func (h *AlbumHandler) UpdateAlbumPrice(c *gin.Context) {
    id := c.Param("id")
    priceStr := c.Query("price")

    price, err := strconv.ParseFloat(priceStr, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price format"})
        return
    }

    if h.AlbumService.UpdateAlbumPrice(price, id) != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update album price"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Successfully updated price"})
}

func (h *AlbumHandler) DeleteAlbumByID(c *gin.Context) {
    id := c.Param("id")

    if h.AlbumService.DeleteAlbumByID(id) != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete album"})
    }

    c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted album"})
}

