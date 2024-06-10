package data

import (
	"github.com/dronestock/deb/internal/internal/config"
)

type Shortcut struct {
	*config.Package
	*config.Source
	*config.To
}
