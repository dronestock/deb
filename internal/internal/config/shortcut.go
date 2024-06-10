package config

type Shortcut struct {
	// 图标
	Icon string `default:"${ICON}" json:"icon,omitempty" validate:"omitempty,required_if=Shortcut true,file"`
	// 是否终端执行
	Terminal bool `default:"${TERMINAL}"`
}
