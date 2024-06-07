package config

type To struct {
	// 安装路径
	Target string `default:"${TARGET=/opt}" json:"target,omitempty"`
	// 架构
	Arch string `default:"${ARCH=amd64}" json:"arch,omitempty" validate:"oneof=i386 amd64 all"`
	// 架构列表
	Architectures []string `default:"${ARCHITECTURES}" json:"architectures,omitempty"`
	// 快捷方式
	Shortcut bool `default:"${SHORTCUT=true}" json:"shortcut,omitempty"`
	// 图标
	Icon string `default:"${ICON}" json:"icon,omitempty" validate:"omitempty,required_if=Shortcut true,file"`
}
