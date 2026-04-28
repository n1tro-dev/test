package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func main() {

	ctx := context.Background()

	conn, err := CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	er := conn.Ping(ctx)
	if er != nil {
		fmt.Println("Ошибка с подключением")
	}
	fmt.Println("Succes")
}

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:postgreAdmin@localhost:5432/testDB")
}
