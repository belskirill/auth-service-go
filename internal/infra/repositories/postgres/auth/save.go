package auth

import (
	"context"
)

func (repo *Repository) SaveUser(ctx context.Context, email string, passHash string) (int64, error) {
	const query = `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id
`
	var id int64
	err := repo.db.QueryRowContext(ctx, query, email, passHash).Scan(&id)
	if err != nil {
		return 0, nil
	}

	return id, nil
}
