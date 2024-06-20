package domain

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/sujeetkumarjha/go-recipes/libs/domain/definitions"
	"github.com/sujeetkumarjha/go-recipes/libs/platform"
	"github.com/sujeetkumarjha/go-recipes/libs/shared_model"
	"net/http"
)

type RecipesHandler struct {
	store definitions.Store
}

func NewRecipesHandler(store definitions.Store) *RecipesHandler {
	return &RecipesHandler{
		store: store,
	}
}

func (h RecipesHandler) CreateRecipe(c *gin.Context) {
	// Get request body and convert it to recipes.Recipe
	var recipe shared_model.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create a url friendly name
	id := slug.Make(recipe.Name)

	// add to the store
	err := h.store.Add(id, recipe)
	if err != nil {
		return
	}

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h RecipesHandler) ListRecipes(c *gin.Context) {
	r, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(200, r)
}

func (h RecipesHandler) GetRecipe(c *gin.Context) {
	id := c.Param("id")

	recipe, err := h.store.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(200, recipe)
}

func (h RecipesHandler) UpdateRecipe(c *gin.Context) {
	// Get request body and convert it to recipes.Recipe
	var recipe shared_model.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	err := h.store.Update(id, recipe)
	if err != nil {
		if errors.Is(err, platform.ErrRecipeNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h RecipesHandler) DeleteRecipe(c *gin.Context) {
	id := c.Param("id")

	err := h.store.Delete(id)
	if err != nil {
		if errors.Is(err, platform.ErrRecipeNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})

}
