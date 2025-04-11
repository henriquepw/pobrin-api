package account

import (
	"time"

	"github.com/henriquepw/pobrin-api/pkg/id"
)

type BudgeRole struct {
	Basic uint8 `json:"basic" db:"budge_basic" validate:"min=0,max=100"`
	Fun   uint8 `json:"fun" db:"budge_fun" validate:"min=0,max=100"`
	Saves uint8 `json:"saves" db:"budge_saves" validate:"min=0,max=100"`
}

type Account struct {
	ID        id.ID     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	BudgeRole BudgeRole `json:"budgeRole"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt time.Time `json:"deletedAt" db:"deleted_at"`
}

type AccountCreate struct {
	Name string `json:"name" validate:"required,min=3"`
}

type AccountUpdate struct {
	Name      *string    `json:"name" validate:"omitempty,min=3"`
	Email     *string    `json:"email" validate:"omitempty,email"`
	BudgeRole *BudgeRole `json:"budgeRole"`
}
