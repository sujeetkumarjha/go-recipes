package definitions

import (
	"github.com/sujeetkumarjha/go-recipes/libs/shared_model"
)

type Store interface {
	Add(name string, recipe shared_model.Recipe) error
	Get(name string) (shared_model.Recipe, error)
	List() (map[string]shared_model.Recipe, error)
	Update(name string, recipe shared_model.Recipe) error
	Delete(name string) error
}
