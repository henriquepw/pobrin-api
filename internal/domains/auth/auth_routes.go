package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/henriquepw/pobrin-api/internal/domains/auth/account"
	"github.com/henriquepw/pobrin-api/internal/domains/auth/credential"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) func(r chi.Router) {
	accountStore := account.NewStore(db)
	accountSVC := account.NewService(accountStore)

	credentialStore := credential.NewStore(db)
	credentialSVC := credential.NewService(credentialStore)

	handler := NewHandler(accountSVC, credentialSVC)

	return func(r chi.Router) {
		r.Post("/sign-up", handler.PostSignUp)
	}
}
