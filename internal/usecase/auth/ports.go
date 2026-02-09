package auth

import (
	"context"
	"os/user"
)

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash string) (int64, error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (user.User, error)
}
