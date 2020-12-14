package platform

import (
	"github.com/hardjonn/geferti/pkg/errs"

	"github.com/denisbrodbeck/machineid"
)

// Host fetches data about the host where app is running on.
type Host struct {
	m func(appID string) (string, error)
}

// NewHost returns a new host
func NewHost() *Host {

	m := machineid.ProtectedID

	return &Host{m: m}
}

// GetMachineID returns the crypto hashed machine id.
func (h *Host) GetMachineID(appKey string) (string, error) {
	id, err := h.m(appKey)

	if err != nil {
		return "", errs.E(errs.Op("platform.repository.GetMachineID"), err, errs.StatusIO)
	}

	return id, nil
}
