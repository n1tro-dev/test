package sqlhand

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRows(ctx context.Context, conn *pgx.Conn, listIDs []int) error {
	sqlQuery := `
	DELETE FROM books
	WHERE id = ANY($1)
	`

	_, err := conn.Exec(ctx, sqlQuery, listIDs)
	if err != nil {
		return err
	}

	return nil
}
