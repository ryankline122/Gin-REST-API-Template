package services

import (
    "example/web-service-gin/internal/models"
    "example/web-service-gin/data"
)

type AlbumService struct {}

func (service *AlbumService) GetAllAlbums() []models.Album {
    return data.Data
}

