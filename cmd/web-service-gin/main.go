package main

import (
	"example/web-service-gin/internal/handlers"
	"example/web-service-gin/internal/routes"
	"example/web-service-gin/internal/services"
)

func main() {
    albumService := &services.AlbumService{}

    albumHandler := &handlers.AlbumHandler{
        AlbumService: albumService,
    }

    router := routes.SetupRouter(albumHandler)
    router.Run("localhost:8080")
}

