package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/vogonwann/gorm-gin/config"
	"github.com/vogonwann/gorm-gin/controller"
	"github.com/vogonwann/gorm-gin/helper"
	"github.com/vogonwann/gorm-gin/model"
	"github.com/vogonwann/gorm-gin/repository"
	"github.com/vogonwann/gorm-gin/router"
	"github.com/vogonwann/gorm-gin/service"
)

func main() {
	log.Info().Msg("Starting server...")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})
	db.Table("users").AutoMigrate(&model.User{})

	// Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	// service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
