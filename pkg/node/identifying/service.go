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

	// AddNode adds a new node to the storage repository.
	AddNode(Node) error
}

type service struct {
	r Repository
}

// NewService creates an identifying service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

// IdentifyNode stands for creating a record in the DB which preserves the node.
func (s *service) IdentifyNode(key string) error {
	// n, err := s.r.GetNodeByKey(key)

	// if err == nil {
	// 	return nil
	// }

	// if err != nil && err != ErrNotFound {
	// 	return errs.E(errs.Op("node.identifying.service.IdentifyNode"), err)
	// }

	// // add Node adding package with the node structure
	// // add Node listing package with node structure

	// // add node here
	// n = Node{}

	return nil
}
