package sqlhand

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn, book Book) error {
	sqlQuery := `
	UPDATE books
	SET name = $1, author = $2, recensia = $3, readed = $4, created_at = $5, readed_at = $6
	WHERE id = $7;
	`
	_, err := conn.Exec(
		ctx,
		sqlQuery,
		book.Name,
		book.Author,
		book.Recensia,
		book.Readed,
		book.CreatedAt,
		book.ReadedAt,
		book.ID,
	)
	if err != nil {
		return err
	}
	return nil
}