package config

import (
	"github.com/dronestock/deb/internal/internal/config/internal"
	"github.com/goexl/gox"
)

type Package struct {
	// 包名
	Name string `default:"${NAME}" json:"name,omitempty"`
	// 版本号
	Version string `default:"${VERSION}" json:"version,omitempty"`
	// 软件的类别
	Section string `default:"${SECTION=utils}" json:"section,omitempty" validate:"oneof=utils net mail text devel"`
	// 软件对于系统的重要程度
	Priority string `default:"${PRIORITY=standard}" json:"priority,omitempty" validate:"oneof=required standard optional extra"`
	// 维护者
	Maintainer *internal.Maintainer `default:"${MAINTAINER}" json:"maintainer,omitempty"`
	// 安装后的大小
	Size gox.Bytes `json:"size,omitempty" validate:"required"`
	// 供应者
	Provide string `default:"${PROVIDE}" validate:"required"`
	// 描述
	Description string `default:"${DESCRIPTION}" json:"description,omitempty"`
	// 申明是否是系统最基本的软件包
	Essential string `default:"${ESSENTIAL=no}" json:"essential,omitempty"`
	// 软件所依赖的其他软件包和库
	// 如果是依赖多个软件包，彼此之间采用英文逗号隔开
	Depends []string `default:"${DEPENDENCIES}" json:"dependencies,omitempty"`
	// 这个字段表明推荐的安装的其他软件包
	Recommends []string `default:"${RECOMMENDS}" json:"recommends,omitempty"`
	// 建议安装的其他软件包
	Suggests []string `default:"${SUGGESTS}" json:"suggests,omitempty"`
}
