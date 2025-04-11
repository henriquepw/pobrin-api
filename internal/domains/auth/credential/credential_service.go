package credential

import (
	"context"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/errors"
	"github.com/henriquepw/pobrin-api/pkg/hash"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/validate"
)

type CredentialService interface {
	CreateCredential(ctx context.Context, dto CredentialCreate) (*Credential, error)
	DeleteCredential(ctx context.Context, id id.ID) error
	GetCredential(ctx context.Context, id id.ID) (*Credential, error)
}

type accouService struct {
	store CredentialStore
}

func NewService(store CredentialStore) CredentialService {
	return &accouService{store}
}

func (s *accouService) CreateCredential(ctx context.Context, dto CredentialCreate) (*Credential, error) {
	if err := validate.Check(dto); err != nil {
		return nil, err
	}

	secret, err := hash.Generate(dto.Password)
	if err != nil {
		return nil, errors.Internal()
	}

	now := time.Now()
	credential := Credential{
		ID:        id.New(),
		User:      dto.User,
		Secret:    secret,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = s.store.Insert(ctx, credential)
	if err != nil {
		return nil, errors.Internal("Failed to create the credential")
	}

	return &credential, nil
}

func (s *accouService) DeleteCredential(ctx context.Context, id id.ID) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		return errors.Internal("Failed to delete the credential")
	}

	return nil
}

func (s *accouService) GetCredential(ctx context.Context, id id.ID) (*Credential, error) {
	item, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, errors.NotFound("Credential not found")
	}

	return item, nil
}
