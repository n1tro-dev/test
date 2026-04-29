package main

import (
	"context"
	"fmt"
	"taskdb/conn"
	"taskdb/sqlhand"

	"github.com/k0kubun/pp"
)

func main() {

	ctx := context.Background()

	conn, err := conn.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	er := conn.Ping(ctx)
	if er != nil {
		fmt.Println("Ошибка с подключением")
	}

	if err := sqlhand.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	// book := sqlhand.Book{
	// 	Name: "Книга5",
	// 	Author: "Human5",
	// 	Recensia: "blaballabl",
	// 	Readed: false,
	// 	CreatedAt: time.Now(),
	// 	ReadedAt: nil,
	// }
	// if err := sqlhand.InsertRow(ctx, conn, book); err != nil {
	// 	panic(err)
	// }

	books, err := sqlhand.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	listIDs:=[]int{1,2,3}
	if err := sqlhand.DeleteRows(ctx, conn, listIDs); err!=nil{
		panic(err)
	}


	// for _, book := range books {
	// 	if book.ID == 1 {
	// 		book.Name = "Собаха"
	// 		book.Author = "НЕзнакомец"
	// 		book.Recensia = "интересный"
	// 		book.Readed = true
	// 		now := time.Now()
	// 		book.ReadedAt = &now
	// 	}

	// 	err := sqlhand.UpdateRow(ctx, conn, book)
	// 	if err != nil{
	// 		panic(err)
	// 	}
	// 	break
	// }
	pp.Println(books)

	fmt.Println("Succes")

}
