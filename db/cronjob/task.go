package cronjob

import (
	"context"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UpdateTaskSummary(ctx context.Context, p *pgxpool.Pool) error {
	s, err := gocron.NewScheduler()
	if err != nil {
		return err
	}
	_, err = s.NewJob(
		gocron.DurationJob(10*time.Second),
		gocron.NewTask(func() {
			tx, err := p.Begin(ctx)
			if err != nil {
				return
			}
			q := `
				UPDATE task_summaries ts
				SET task_count = t.cnt, updated_at = NOW()
				FROM (
					SELECT user_id, COUNT(*) AS cnt
					FROM tasks
					GROUP BY user_id
				) t
				WHERE ts.user_id = t.user_id;
			`
			_, err = tx.Exec(ctx, q)
			if err != nil {
				tx.Rollback(ctx)

				return
			}
			tx.Commit(ctx)
		}),
	)
	s.Start()

	return nil
}
