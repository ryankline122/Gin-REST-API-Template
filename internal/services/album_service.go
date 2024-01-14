package services

import (
	"database/sql"
	"errors"
	"example/web-service-gin/internal/models"
)

type AlbumService struct {
    DatabaseConnection *sql.DB
}

// Create
func (s *AlbumService) AddAlbum(album models.Album) error {
    query := `
        INSERT INTO album (title, artist, price)
        VALUES ($1, $2, $3);
    `

    _, err := s.DatabaseConnection.Exec(query, album.Title, album.Artist, album.Price)
    if err != nil {
       return err 
    }

    return nil
}

// Read
func (s *AlbumService) GetAllAlbums() ([]models.Album, error) {
    rows, err := s.DatabaseConnection.Query(`SELECT * FROM album;`)
    if err != nil {
        return []models.Album{}, err 
    }
    defer rows.Close()

    albums := make([]models.Album, 0)
    for rows.Next() {
        var album models.Album
        if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
            return []models.Album{}, err 
        }
        albums = append(albums, album)
    }

    // Check for errors during iteration
    if err := rows.Err(); err != nil {
        return []models.Album{}, err
    }

    if len(albums) == 0 {
        return albums, errors.New("No albums found")
    }

    return albums, nil
}

func (s *AlbumService) GetAlbumByID(id string) (models.Album, error) {
    var album models.Album
    query := `
        SELECT * FROM album
        WHERE id = $1;
    `

    row := s.DatabaseConnection.QueryRow(query, id)
    if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
        return models.Album{}, err 
    }

    return album, nil
}

// Update
func (s *AlbumService) UpdateAlbumPrice(price float64, id string) error {
    query := `
        UPDATE album
        SET price = $1
        WHERE id = $2;
    `

    _, err := s.DatabaseConnection.Exec(query, price, id)
    if err != nil {
       return err 
    }

    return nil
}

// Delete
func (s *AlbumService) DeleteAlbumByID(id string) error {
    query := `
        DELETE FROM album
        WHERE id = $1;
    `

    _, err := s.DatabaseConnection.Exec(query, id)
    if err != nil {
        return err
    }

    return nil
}

