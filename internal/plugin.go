package internal

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/todo/internal/internal/step"
)

type plugin struct {
	drone.Base
}

func New() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewTodo()).Name("样例").Build(),
	}
}
