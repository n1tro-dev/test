package sqlhand

import (
	"context"

	"github.com/jackc/pgx/v5"
)

//	func InsertRow(ctx context.Context, conn *pgx.Conn, name string, author string, recensia string, readed bool, created_at time.Time) error {
//		sqlQuery := `
//		INSERT INTO books (name, author, recensia, readed, created_at)
//		VALUES($1, $2, $3, $4, $5);
//		`
//		_, err := conn.Exec(ctx, sqlQuery, name, author, recensia, readed, created_at)
//		if err != nil {
//			return err
//		}
//		return nil
//	}
func InsertRow(ctx context.Context, conn *pgx.Conn, book Book) error {
	sqlQuery := `
	INSERT INTO books (name, author, recensia, readed, created_at)
	VALUES($1, $2, $3, $4, $5);
	`
	_, err := conn.Exec(ctx, sqlQuery, book.Name, book.Author, book.Recensia, book.Readed, book.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
