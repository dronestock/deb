package config

type Package struct {
	// 包名
	Name string `default:"${NAME}" json:"name,omitempty"`
	// 版本号
	Version string `default:"${VERSION}" json:"version,omitempty"`
	// 软件的类别
	Section string `default:"${SECTION=utils}" json:"section,omitempty" validate:"oneof=utils net mail text devel"`
	// 软件对于系统的重要程度
	Priority string `default:"${PRIORITY=standard}" json:"priority,omitempty" validate:"oneof=required standard optional extra"`
	// 架构
	Arch string `default:"${ARCH=amd64}" json:"arch,omitempty" validate:"oneof=i386 amd64 all"`
	// 维护者
	Maintainer Maintainer `default:"${MAINTAINER}" json:"maintainer,omitempty"`
	// 供应者
	Provide string `default:"${PROVIDE}" validate:"required"`
	// 描述
	Description string `default:"${DESCRIPTION}" json:"description,omitempty"`
}
