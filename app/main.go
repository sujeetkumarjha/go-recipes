package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sujeetkumarjha/go-recipes/libs/domain"
	"github.com/sujeetkumarjha/go-recipes/libs/platform"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Recipes API
func main() {
	//Create Gin router
	router := gin.Default()
	//add swagger

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Instantiate recipe Handler and provide a data store
	store := platform.NewInMemoryStore()
	handler := domain.NewRecipesHandler(store)

	// Define routes
	router.GET("/recipes", handler.ListRecipes)
	router.GET("/recipes/:id", handler.GetRecipe)
	router.POST("/recipes", handler.CreateRecipe)
	router.PUT("/recipes/:id", handler.UpdateRecipe)
	router.DELETE("/recipes/:id", handler.DeleteRecipe)

	// Start the server
	router.Run()
}
