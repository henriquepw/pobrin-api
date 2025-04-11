package auth

import (
	"net/http"

	"github.com/henriquepw/pobrin-api/internal/domains/auth/account"
	"github.com/henriquepw/pobrin-api/internal/domains/auth/credential"
	"github.com/henriquepw/pobrin-api/pkg/httputil"
)

type authHandler struct {
	account    account.AccountService
	credential credential.CredentialService
}

func NewHandler(a account.AccountService, c credential.CredentialService) *authHandler {
	return &authHandler{a, c}
}

func (h *authHandler) PostSignUp(w http.ResponseWriter, r *http.Request) {
	body, err := httputil.GetBodyRequest[SignUpRequest](r)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	account, err := h.account.CreateAccount(r.Context(), account.AccountCreate{
		Name: body.Name,
	})
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	_, err = h.credential.CreateCredential(r.Context(), credential.CredentialCreate{
		User:     body.Email,
		Password: body.Password,
	})
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessCreatedResponse(w, account.ID.String())
}
