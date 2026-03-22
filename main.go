package main

import (
	"fmt"
	"study/feature1"
	"study/feature2"
	simpleconnection "study/feture_postgres/simple_connection"
)

func main() {
	fmt.Println("Это главный файл")
	feature1.Feature1()

	feature2.Feature2()
	simpleconnection.CheckConnection()
}