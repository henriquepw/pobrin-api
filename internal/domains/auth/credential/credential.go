package credential

import (
	"time"

	"github.com/henriquepw/pobrin-api/pkg/id"
)

type Credential struct {
	ID        id.ID     `json:"id" db:"id"`
	User      string    `json:"user" db:"user"`
	Secret    string    `json:"secret" db:"secret"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt time.Time `json:"deletedAt" db:"deleted_at"`
}

type CredentialCreate struct {
	User     string `json:"user" validate:"required,min=3"`
	Password string `json:"password" validate:"required,email"`
}
