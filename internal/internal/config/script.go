package config

type Script struct {
	Pre  string `json:"pre,omitempty" validate:"omitempty,file"`
	Post string `json:"post,omitempty" validate:"omitempty,file"`
}
