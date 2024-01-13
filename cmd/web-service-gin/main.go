package main

import (
	"example/web-service-gin/data"
	"example/web-service-gin/internal/handlers"
	"example/web-service-gin/internal/routes"
	"example/web-service-gin/internal/services"
)

func main() {
    dbConnection, err := data.ConnectToDatabase()
    if err != nil {
        panic(err)
    }
    defer dbConnection.Close()

    albumService := &services.AlbumService{
        DatabaseConnection: dbConnection,
    }

    albumHandler := &handlers.AlbumHandler{
        AlbumService: albumService,
    }

    router := routes.SetupRouter(albumHandler)
    router.Run("localhost:8080")
}

