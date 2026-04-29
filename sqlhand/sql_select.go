package sqlhand

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]Book, error) {
	sqlQuery := `
	SELECT id, name, author, recensia, readed, created_at, readed_at
	FROM books
	ORDER BY id ASC;
	`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := make([]Book, 0)

	for rows.Next(){
		var book Book

		err := rows.Scan(
			&book.ID,
			&book.Name,
			&book.Author,
			&book.Recensia,
			&book.Readed,
			&book.CreatedAt,
			&book.ReadedAt,
		)
		if err != nil{
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}