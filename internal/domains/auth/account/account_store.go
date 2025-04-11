package account

import (
	"context"
	"database/sql"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/date"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/jmoiron/sqlx"
)

type AccountStore interface {
	Insert(ctx context.Context, i Account) error
	Delete(ctx context.Context, id id.ID) error
	Update(ctx context.Context, id id.ID, dto AccountUpdate) error
	Get(ctx context.Context, id id.ID) (*Account, error)
}

type accountStore struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func NewStore(db *sqlx.DB) AccountStore {
	return &accountStore{db, nil}
}

func (s *accountStore) StartTransaction() error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	s.tx = tx
	return nil
}

func (s *accountStore) SetTransaction(tx *sqlx.Tx) {
	s.tx = tx
}

func (s *accountStore) CloseTransaction() {
	s.tx = nil
}

type DB interface {
	NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	GetContext(ctx context.Context, dest any, query string, args ...any) error
}

func (s *accountStore) getDB() DB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

func (s *accountStore) Insert(ctx context.Context, i Account) error {
	query := `
    INSERT INTO accounts (id, name, email, budge_basic, budge_fun, budge_saves, created_at, updated_at)
		VALUES (:id, :name, :email, :budge_basic, :budge_fun, :budge_saves, :created_at, :updated_at)
  `
	_, err := s.getDB().NamedExecContext(ctx, query, i)

	return err
}

func (s *accountStore) Delete(ctx context.Context, id id.ID) error {
	_, err := s.getDB().ExecContext(ctx, "DELETE FROM accounts WHERE id = ?", id)
	return err
}

func (s *accountStore) Update(ctx context.Context, id id.ID, i AccountUpdate) error {
	query := `
    UPDATE accounts
    SET name = ?, email = ?, updated_at = ?
    WHERE id = ?
  `
	_, err := s.getDB().ExecContext(
		ctx, query,
		i.Name,
		i.Email,
		date.FormatToISO(time.Now()),
		id,
	)

	return err
}

func (s *accountStore) Get(ctx context.Context, id id.ID) (*Account, error) {
	query := "SELECT * FROM accounts WHERE id = ?"

	var item Account
	err := s.getDB().GetContext(ctx, &item, query, id)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
