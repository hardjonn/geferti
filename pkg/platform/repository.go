package platform

import (
	"github.com/denisbrodbeck/machineid"
)

// Host fetches data about the host where app is running on.
type Host struct {
	m func(appID string) (string, error)
}

// NewHost return a new host
func NewHost() *Host {

	m := machineid.ProtectedID

	return &Host{m: m}
}

// GetMachineID returns the crypto hashed machine id.
func (h *Host) GetMachineID(appKey string) (string, error) {
	id, err := h.m(appKey)

	if err != nil {
		return "", err
	}

	return id, nil
}
