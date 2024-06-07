package internal

import (
	"github.com/dronestock/deb/internal/internal/config"
	"github.com/dronestock/deb/internal/internal/step"
	"github.com/dronestock/drone"
)

type plugin struct {
	drone.Base

	Package config.Package `default:"${PACKAGE}" json:"package,omitempty"`
	Source  config.Source  `default:"${SOURCE}" json:"source,omitempty"`
	To      config.To      `default:"${TO}" json:"to,omitempty"`
}

func New() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewPrepare()).Name("样例").Build(),
	}
}
