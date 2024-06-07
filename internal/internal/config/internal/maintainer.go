package internal

import (
	"fmt"
)

type Maintainer struct {
	Name  string `json:"name,omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required,email"`
}

func (m *Maintainer) String() string {
	return fmt.Sprintf("%s <%s>", m.Name, m.Email)
}
