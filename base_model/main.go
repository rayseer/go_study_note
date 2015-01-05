package main

import (
	"./model"
	"fmt"
)

func main() {
	fmt.Println("Start...")
	db = model.Connect()
	fmt.Println(db)
}
