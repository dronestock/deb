package step

import (
	"context"
	_ "embed"
	"fmt"
	"sync"

	"github.com/dronestock/deb/internal/internal/config"
	"github.com/dronestock/deb/internal/internal/step/internal/data"
	"github.com/goexl/gox/template"
)

//go:embed internal/template/control.gohtml
var controlTemplate string

type Prepare struct {
	pack   *config.Package
	source *config.Source
	to     *config.To
}

func NewPrepare() *Prepare {
	return &Prepare{
		// 字段初步化
	}
}

func (p *Prepare) Runnable() bool {
	return true
}

func (p *Prepare) Run(ctx context.Context) (err error) {
	architectures := []string{
		p.to.Arch,
	}
	architectures = append(architectures, p.to.Architectures...)

	wg := new(sync.WaitGroup)
	wg.Add(len(architectures))
	for _, architecture := range architectures {

	}

	return
}

func (p *Prepare) run(ctx context.Context, arch string, err *error) {
	work := fmt.Sprintf("/tmp/deb/%s", arch)

	return
}

func (p *Prepare) makeControl(_ context.Context, arch string) error {
	data := new(data.Control)
	template.New(controlTemplate).Data(d.config).Build().ToFile(d.config.Filepath())
}
