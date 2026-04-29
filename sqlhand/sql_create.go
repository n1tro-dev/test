package sqlhand

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS books (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	author VARCHAR(100) NOT NULL,
	recensia VARCHAR(1000),
	readed BOOLEAN NOT NULL,

	created_at TIMESTAMP NOT NULL,
	readed_at TIMESTAMP
	);
	`

	_, err := conn.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}
