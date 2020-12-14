package zerolog

import (
	"github.com/hardjonn/geferti/pkg/config"
	"github.com/hardjonn/geferti/pkg/logging"
)

// Factory is a receiver for zerolog factory
type Factory struct{}

// Build zerolog logger
func (mf *Factory) Build(lc *config.Logger) (logging.Logger, error) {
	l, err := New(lc)
	if err != nil {
		return nil, err
	}

	return l, nil
}
