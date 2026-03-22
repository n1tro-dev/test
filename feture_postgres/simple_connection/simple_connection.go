package simpleconnection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgreAdmin@localhost:5432/gocourse")

	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Успешно поключено")
	

}