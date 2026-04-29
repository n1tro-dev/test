package main

import (
	"context"
	"fmt"
	simpleconnection "study/feture_postgres/simple_connection"
	simplesql "study/feture_postgres/simple_sql"
)

func main() {
	ctx := context.Background()

	conn, err := simpleconnection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simplesql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	// if err := simplesql.InserRow(ctx, conn, "Домашка", "сделать дз", false, time.Now()); err != nil {
	// 	panic(err)
	// }

	// if err := simplesql.DeleteRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// tasks, err := simplesql.SelectRows(ctx, conn)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, task := range tasks {
	// 	if task.ID == 3 {
	// 		task.Title = "Покормить кошку"
	// 		task.Description = "Насыпать кошке корм"
	// 		task.Completed = true
	// 		now := time.Now()
	// 		task.CompletedAt = &now

	// 		err := simplesql.UpdateTask(ctx, conn, task)
	// 		if err != nil{
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// }

	// pp.Println(tasks)

	fmt.Println("Успешно")
}
