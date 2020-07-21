package hosting

// Service provides host specific operations
type Service interface {
	GetMachineID(string) (string, error)
}

// Repository provides access to the platform layer
type Repository interface {
	// GetMachineID returns a crypto hashed machine ID
	GetMachineID(string) (string, error)
}

type service struct {
	r Repository
}

// NewService creates a host service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetMachineID(appKey string) (string, error) {
	key, err := s.r.GetMachineID(appKey)

	if err != nil {
		return "", err
	}

	return key, nil
}
