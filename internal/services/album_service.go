package services

import (
	"errors"
	"example/web-service-gin/data"
	"example/web-service-gin/internal/models"
	"fmt"
)

type AlbumService struct {}

// Create
func (s *AlbumService) AddAlbum(album models.Album) {
    data.Data = append(data.Data, album)
}

// Read
func (s *AlbumService) GetAllAlbums() []models.Album {
    return data.Data
}

func (s *AlbumService) GetAlbumByID(id string) (models.Album, error) {
    for _, album := range data.Data {
        if album.ID == id {
            return album, nil
        }
    }

    return models.Album{}, errors.New(fmt.Sprintf("Album with ID %s not found", id))
}

// Update
func (s *AlbumService) UpdateAlbumPrice(price float64, id string) error {
    for i, album := range data.Data {
        if album.ID == id {
            data.Data[i].Price = price
            return nil
        }
    }

    return errors.New(fmt.Sprintf("Album with ID %s not found", id))

}

// Delete
func (s *AlbumService) DeleteAlbumByID(id string) error {
    for i, album := range data.Data {
        if album.ID == id {
            copy(data.Data[i:], data.Data[i+1:])
            data.Data = data.Data[:len(data.Data)-1]
            return nil
        }
    }

    return errors.New(fmt.Sprintf("Album with ID %s not found", id))
}

