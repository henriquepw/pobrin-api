package credential

import (
	"context"

	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/jmoiron/sqlx"
)

type CredentialStore interface {
	Insert(ctx context.Context, i Credential) error
	Delete(ctx context.Context, id id.ID) error
	Get(ctx context.Context, id id.ID) (*Credential, error)
}

type credentialStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) CredentialStore {
	return &credentialStore{db}
}

func (s *credentialStore) Insert(ctx context.Context, i Credential) error {
	query := `
    INSERT INTO credentials (id, name, secret, created_at, updated_at)
		VALUES (:id, :name, :secret, :created_at, :updated_at)
  `
	_, err := s.db.NamedExecContext(ctx, query, i)

	return err
}

func (s *credentialStore) Delete(ctx context.Context, id id.ID) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM credentials WHERE id = ?", id)
	return err
}

func (s *credentialStore) Get(ctx context.Context, id id.ID) (*Credential, error) {
	query := "SELECT * FROM credentials WHERE id = ?"

	var item Credential
	err := s.db.GetContext(ctx, &item, query, id)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
