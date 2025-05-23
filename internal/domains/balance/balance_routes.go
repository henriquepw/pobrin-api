package balance

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) func(r chi.Router) {
	store := NewStore(db)
	svc := NewService(store)
	handler := NewHandler(svc)

	return func(r chi.Router) {
		r.Get("/", handler.GetUserBalance)
		r.Put("/", handler.PutUserBalance)
	}
}
