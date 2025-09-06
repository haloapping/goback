package task

import (
	"fmt"
	"strings"

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

func (r *Repository) Add(c echo.Context, req AddReq) (Task, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Task{}, err
	}

	q := `
		INSERT INTO tasks(id, user_id, title, description)
		VALUES($1, $2, $3, $4)
		RETURNING *;
	`
	row := tx.QueryRow(ctx, q, ulid.Make().String(), req.UserId, req.Title, req.Description)
	var t Task
	err = row.Scan(&t.Id, &t.UserId, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return Task{}, err
	}
	tx.Commit(ctx)

	return t, nil
}

func (r *Repository) FindById(c echo.Context, id string) (Task, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Task{}, err
	}

	q := `
		SELECT * FROM tasks
		WHERE id = $1;
	`
	row := tx.QueryRow(ctx, q, id)
	var t Task
	err = row.Scan(&t.Id, &t.UserId, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return Task{}, err
	}
	tx.Commit(ctx)

	return t, nil
}

func (r *Repository) FindAllTasksByUserId(c echo.Context, id string) ([]UserTask, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return []UserTask{}, err
	}

	q := `
		SELECT COUNT(*) FROM tasks t 
		LEFT JOIN users u ON t.user_id = u.id
		WHERE u.id = $1;
	`
	row := tx.QueryRow(ctx, q, id)
	var count int
	err = row.Scan(&count)
	if err != nil {
		tx.Rollback(ctx)

		return []UserTask{}, err
	}

	q = `
		SELECT t.id, t.title, t.description, t.created_at, t.updated_at FROM tasks t 
		LEFT JOIN users u ON t.user_id = u.id
		WHERE u.id = $1;
	`
	rows, err := tx.Query(ctx, q, id)
	rowNum := 0
	userTasks := make([]UserTask, count)
	for rows.Next() {
		var t UserTask
		rows.Scan(&t.Id, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			tx.Rollback(ctx)

			return []UserTask{}, err
		}
		userTasks[rowNum] = t
		rowNum++
	}
	tx.Commit(ctx)

	return userTasks, nil
}

func (r *Repository) FindAll(c echo.Context) ([]Task, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return []Task{}, err
	}

	q := `
		SELECT COUNT(*) FROM tasks;
	`
	row := tx.QueryRow(ctx, q)
	var count int
	err = row.Scan(&count)
	if err != nil {
		tx.Rollback(ctx)

		return []Task{}, err
	}

	q = `
		SELECT * FROM tasks;
	`
	rows, err := tx.Query(ctx, q)
	rowNum := 0
	tasks := make([]Task, count)
	for rows.Next() {
		var t Task
		rows.Scan(&t.Id, &t.UserId, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			tx.Rollback(ctx)

			return []Task{}, err
		}
		tasks[rowNum] = t
		rowNum++
	}
	tx.Commit(ctx)

	return tasks, nil
}

func (r *Repository) UpdateById(c echo.Context, id string, req UpdateReq) (Task, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Task{}, err
	}

	columnNames := make([]string, 0)
	columnValues := make([]interface{}, 0)
	argPos := 1
	if req.Title != "" {
		columnNames = append(columnNames, fmt.Sprintf("title = $%d", argPos))
		columnValues = append(columnValues, req.Title)
		argPos++
	}
	if req.Description != "" {
		columnNames = append(columnNames, fmt.Sprintf("description = $%d", argPos))
		columnValues = append(columnValues, req.Description)
		argPos++
	}
	columnValues = append(columnValues, id)

	q := fmt.Sprintf(`
		UPDATE tasks
		SET %s
		WHERE id = $%d
		RETURNING *;
	`, strings.Join(columnNames, ", "), argPos)
	row := tx.QueryRow(ctx, q, columnValues...)
	var t Task
	err = row.Scan(&t.Id, &t.UserId, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return Task{}, err
	}
	tx.Commit(ctx)

	return t, nil
}

func (r *Repository) DeleteById(c echo.Context, id string) (Task, error) {
	ctx := c.Request().Context()
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)

		return Task{}, err
	}

	q := `
		DELETE FROM tasks
		WHERE id = $1
		RETURNING *;
	`
	var t Task
	row := tx.QueryRow(ctx, q, id)
	err = row.Scan(&t.Id, &t.UserId, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return Task{}, err
	}
	tx.Commit(ctx)

	return t, nil
}
