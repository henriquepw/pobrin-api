package account

import (
	"context"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/errors"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/validate"
)

type AccountService interface {
	CreateAccount(ctx context.Context, dto AccountCreate) (*Account, error)
	UpdateAccount(ctx context.Context, id id.ID, dto AccountUpdate) error
	DeleteAccount(ctx context.Context, id id.ID) error
	GetAccount(ctx context.Context, id id.ID) (*Account, error)
}

type accouService struct {
	store AccountStore
}

func NewService(store AccountStore) AccountService {
	return &accouService{store}
}

func (s *accouService) CreateAccount(ctx context.Context, dto AccountCreate) (*Account, error) {
	if err := validate.Check(dto); err != nil {
		return nil, err
	}

	now := time.Now()
	account := Account{
		ID:   id.New(),
		Name: dto.Name,
		BudgeRole: BudgeRole{
			Basic: 50,
			Fun:   30,
			Saves: 20,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.store.Insert(ctx, account)
	if err != nil {
		return nil, errors.Internal("Failed to create the account")
	}

	return &account, nil
}

func (s *accouService) UpdateAccount(ctx context.Context, id id.ID, dto AccountUpdate) error {
	if err := validate.Check(dto); err != nil {
		return err
	}

	// TODO:

	return nil
}

func (s *accouService) DeleteAccount(ctx context.Context, id id.ID) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		return errors.Internal("Failed to delete the account")
	}

	return nil
}

func (s *accouService) GetAccount(ctx context.Context, id id.ID) (*Account, error) {
	item, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, errors.NotFound("Account not found")
	}

	return item, nil
}
