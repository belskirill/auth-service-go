package auth

import (
	"context"
	"os/user"
)

func (repo *Repository) User(ctx context.Context, email string) (user.User, error) {
	return user.User{}, nil
}
