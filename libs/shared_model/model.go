package shared_model

// Recipe represents a recipe
type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

// Ingredient represents individual ingredients
type Ingredient struct {
	Name string `json:"name"`
}
