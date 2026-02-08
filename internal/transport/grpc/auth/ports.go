package auth

import (
	"context"
	"proto-auth/internal/domain/auth"
)

type Auth interface {
	Login(ctx context.Context, login auth.Login) error
}
