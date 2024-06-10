package step

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/dronestock/deb/internal/internal/command"
	"github.com/dronestock/deb/internal/internal/config"
	"github.com/dronestock/deb/internal/internal/step/internal/data"
	"github.com/goexl/gfx"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/template"
	"github.com/goexl/log"
)

//go:embed internal/template/control.gohtml
var controlTemplate string

type Prepare struct {
	pkg    *config.Package
	source *config.Source
	to     *config.To

	cp *command.Copy
	logger log.Logger
	_      gox.CannotCopy
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
	work := fmt.Sprintf("/tmp/deb/%s/%s",p.pkg.Name, arch)
	debian :=filepath.Join(work, "DEBIAN")
	source:=filepath.Join(work, p.to.Target)
	if mce := p.mkdir(debian); nil!= mce {
		*err = mce
	}else if mse:=p.mkdir(source);nil!=mse{
		*err=mse
	}else if cse:=p.cp.Exec(ctx, p.source.Directory, source);nil!=cse{
		*err=cse
	}else {
		*err=p.makeControl(ctx, arch,filepath.Join(debian, "control"))
	}

	return
}

func (p *Prepare) processShortcut() (err error) {
	if p.to.Shortcut{
		return
	}



	return
}

func (p *Prepare) makeControl(_ context.Context, arch string,path string) (err error) {
	control := new(data.Control)
	control.Package = p.pkg
	control.Architecture = arch
	err = template.New(controlTemplate).Data(control).Build().ToFile(path))

	return
}

func (p *Prepare) mkdir(path string) (err error) {
	dir := filepath.Dir(path)
	if _, exists := gfx.Exists(dir); !exists {
		err = os.MkdirAll(dir, os.ModePerm)
	}

	fields := gox.Fields[any]{
		field.New("dir", dir),
	}
	if nil != err {
		p.logger.Warn("创建目录出错", fields.Add(field.Error(err))...)
	} else {
		p.logger.Debug("创建目录成功", fields...)
	}

	return
}
