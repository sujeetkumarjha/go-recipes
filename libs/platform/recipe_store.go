package platform

import (
	"errors"
	"github.com/sujeetkumarjha/go-recipes/libs/shared_model"
)

type InMemoryStore struct {
	list map[string]shared_model.Recipe
}

var (
	ErrRecipeNotFound = errors.New("not found")
)

func (s *InMemoryStore) Add(name string, recipe shared_model.Recipe) error {
	s.list[name] = recipe
	return nil
}

func (s *InMemoryStore) Get(name string) (shared_model.Recipe, error) {
	recipe, ok := s.list[name]
	if !ok {
		return shared_model.Recipe{}, ErrRecipeNotFound
	}
	return recipe, nil
}

func (s *InMemoryStore) List() (map[string]shared_model.Recipe, error) {
	return s.list, nil
}

func (s *InMemoryStore) Update(name string, recipe shared_model.Recipe) error {
	if _, ok := s.list[name]; ok {
		s.list[name] = recipe
		return nil
	}

	return ErrRecipeNotFound
}

func (s *InMemoryStore) Delete(name string) error {
	delete(s.list, name)
	return nil
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		list: make(map[string]shared_model.Recipe),
	}
}
