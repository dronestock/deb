package config

type Source struct {
	Execute   string `default:"${EXECUTE=${EXEC}}" json:"execute,omitempty" validate:"required,file"`
	Service   string `default:"${SERVICE}" json:"service,omitempty" validate:"omitempty,file"`
	Directory string `default:"${DIRECTORY=${DIR}}" json:"directory,omitempty" validate:"omitempty,file"`
}
