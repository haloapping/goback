package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type Repository struct {
	Pool *pgxpool.Pool
}

func NewRepository(p *pgxpool.Pool) Repository {
	return Repository{
		Pool: p,
	}
}

func (r *Repository) Register(c echo.Context, req UserRegisterReq) (UserRegister, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return UserRegister{}, err
	}

	q := `
		INSERT INTO users(id, username, password)
		VALUES($1, $2, $3)
		RETURNING id, username, created_at, updated_at;
	`
	row := tx.QueryRow(ctx, q, ulid.Make().String(), req.Username, req.Password)
	var u UserRegister
	err = row.Scan(&u.Id, &u.Username, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return UserRegister{}, err
	}
	tx.Commit(ctx)

	return u, nil
}

func (r *Repository) Login(c echo.Context, req UserLoginReq) (string, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return "", err
	}

	q := `
		SELECT id FROM users
		WHERE username = $1;
	`
	row := tx.QueryRow(ctx, q, req.Username)
	var userId string
	err = row.Scan(&userId)
	if err != nil {
		tx.Rollback(ctx)

		return "", err
	}
	tx.Commit(ctx)

	return userId, nil
}

func (r *Repository) Biodata(c echo.Context, id string) (UserBiodata, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return UserBiodata{}, err
	}

	q := `
		SELECT id, username, created_at, updated_at FROM users
		WHERE id = $1;
	`
	row := tx.QueryRow(ctx, q, id)
	var b UserBiodata
	err = row.Scan(&b.Id, &b.Username, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return UserBiodata{}, err
	}
	tx.Commit(ctx)

	return b, nil
}
