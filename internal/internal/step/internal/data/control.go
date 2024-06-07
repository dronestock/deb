package data

import (
	"github.com/dronestock/deb/internal/internal/config"
)

type Control struct {
	config.Package

	Architecture string `json:"architecture,omitempty"`
}
