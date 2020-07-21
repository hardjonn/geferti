package identifying

import (
	"errors"
)

// ErrNotFound is used when a node could not be found
var ErrNotFound = errors.New("node not found")

// Service provides node identifying operations
type Service interface {
	IdentifyNode(string) error
}

// Repository provides access to node repository
type Repository interface {
	// GetNodeByKey returns a node by the given key.
	GetNodeByKey(string) (Node, error)
}

type service struct {
	r Repository
}

// NewService creates an identifying service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) IdentifyNode(key string) error {
	// n, err := s.r.GetNodeByKey(key)
	return nil
}
